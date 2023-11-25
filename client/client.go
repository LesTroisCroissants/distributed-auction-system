package main

import (
	"context"
	"distributed-auction-system/proto"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	server      proto.AuctionClient
	serverPorts = [3]int{8080, 8081, 8082}
)

func main() {
	connectToServer(serverPorts[0])

	primary, _ := server.WhoIsNewLeader(context.Background(), &proto.Empty{}) // timeout not implemented

	connectToServer(int(primary.Port))
	server.Bid(context.Background(), &proto.ClientBid{Bid: 1})
}

type Server struct {
	proto.UnimplementedAuctionServer
}

func connectToServer(port int) {
	conn, _ := grpc.Dial(":"+strconv.Itoa(port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	server = proto.NewAuctionClient(conn)

	fmt.Println("Connected to port: ", strconv.Itoa(port))
}
