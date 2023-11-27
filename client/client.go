package main

import (
	"context"
	"distributed-auction-system/proto"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	server      proto.AuctionClient
	serverPorts = [3]int{8080, 8081, 8082}
)

type Server struct {
	proto.UnimplementedAuctionServer
}

func main() {
	connectToServer(serverPorts[0])

	for {
		var userInput string
		fmt.Scanln(&userInput)

		//check if server is online

		if strings.ToLower(userInput) == "status" {
			fmt.Println(checkStatus())
		} else {
			bid, _ := strconv.Atoi(userInput)
			makeBid(bid)
		}
	}
}

func findServer(port int) {
	fmt.Println("Connection has been lost, finding new server")
	connectToServer(port)
	primary, _ := server.WhoIsNewLeader(context.Background(), &proto.Empty{}) // timeout not implemented
	fmt.Println("Found the new primary!")
	connectToServer(int(primary.Port))
}

func getErrorMsg(err string) string {
	text := strings.Split(err, " = ")
	return text[2]
}

func makeBid(amount int) {
	_, err := server.Bid(context.Background(), &proto.ClientBid{Bid: int32(amount)})

	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			findServer(serverPorts[1])
			makeBid(amount)
		} else {
			fmt.Println("Server:", getErrorMsg(err.Error()))
		}
	} else {
		fmt.Println("Bid registered!")
	}
}

func checkStatus() string {
	result, err := server.Result(context.Background(), &proto.Empty{})

	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			findServer(serverPorts[1])
		} else {
			fmt.Println("Server:", getErrorMsg(err.Error()))
		}
	}

	return result.Status
}

func connectToServer(port int) {
	conn, _ := grpc.Dial(":"+strconv.Itoa(port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	server = proto.NewAuctionClient(conn)

	fmt.Println("Connected to port: ", strconv.Itoa(port))
}
