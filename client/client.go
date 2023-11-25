package main

import (
	"context"
	"distributed-auction-system/proto"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	server      proto.AuctionClient
	serverPorts = [3]int{8080, 8081, 8082}
	timeout     = 2 * time.Second
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
	connectToServer(port)
	primary, _ := server.WhoIsNewLeader(context.Background(), &proto.Empty{}) // timeout not implemented
	connectToServer(int(primary.Port))
}

func deadline() time.Time {
	return time.Now().Add(timeout)
}

func makeBid(amount int) {
	deadlineContext, cancel := context.WithDeadline(context.Background(), deadline())
	defer cancel()
	_, err := server.Bid(deadlineContext, &proto.ClientBid{Bid: int32(amount)})

	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			findServer(serverPorts[1])
		} else {
			fmt.Println(err.Error())
		}
	}

	if deadlineContext.Err() != nil {
		fmt.Println("Server has crashed, finding new server")
		findServer(serverPorts[1])
	}
}

func checkStatus() string {
	deadlineContext, cancel := context.WithDeadline(context.Background(), deadline())
	defer cancel()
	result, _ := server.Result(deadlineContext, &proto.Empty{})

	if deadlineContext.Err() != nil {
		fmt.Println("Server has crashed, finding new server")
		findServer(serverPorts[1])
	}

	return result.Status
}

func connectToServer(port int) {
	conn, _ := grpc.Dial(":"+strconv.Itoa(port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	server = proto.NewAuctionClient(conn)

	fmt.Println("Connected to port: ", strconv.Itoa(port))
}
