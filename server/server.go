package main

import (
	"distributed-auction-system/proto"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverPorts       = [3]int{8080, 8081, 8082}
	serverConnections = make([]proto.AuctionClient, 0)
	port              = flag.Int("port", 8080, "port number")
)

type Server struct {
	proto.UnimplementedAuctionServer
	port int
}

func main() {
	flag.Parse()

	server := &Server{
		port: *port,
	}

	go startServer(server)

	for _, p := range serverPorts {
		if server.port != p {
			connectToServer(p)
		}
	}

	time.Sleep(time.Hour)
}

func connectToServer(port int) {
	conn, _ := grpc.Dial(":"+strconv.Itoa(port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	replica := proto.NewAuctionClient(conn)
	serverConnections = append(serverConnections, replica)

	fmt.Println("Connected to port: ", strconv.Itoa(port))
}

// code adapted from TAs
// https://github.com/Mai-Sigurd/grpcTimeRequestExample?tab=readme-ov-file#setting-up-the-files
func startServer(server *Server) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterAuctionServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}
