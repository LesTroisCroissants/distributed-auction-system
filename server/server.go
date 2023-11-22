package main

import (
	"context"
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
	serverPorts = [3]int{8080, 8081, 8082}
	servers     = make(map[int]proto.AuctionClient) // map port to connections
	serverId    = flag.Int("id", 1, "port number")
	timeout     = 2 * time.Second //random assumption
	leader      = 8080
	bid         int32
)

type Server struct {
	proto.UnimplementedAuctionServer
	port int
}

func main() {
	flag.Parse()

	server := &Server{
		port: serverPorts[*serverId-1],
	}

	go startServer(server)

	for _, p := range serverPorts {
		if server.port != p {
			connectToServer(p)
		}
	}

	time.Sleep(time.Hour)
}

func (server *Server) Election(ctx context.Context, currentCandidate *proto.RingLeaderTopDawgG) (*proto.Acknowledgement, error) {
	//check for own identifier
	if currentCandidate.ProcessID == int32(server.port) {
		AnnounceResult(currentCandidate)
		return &proto.Acknowledgement{}, nil
	}

	//compare with own bid and identifier
	if currentCandidate.Bid < bid || (currentCandidate.Bid == bid && currentCandidate.ProcessID < int32(server.port)) {
		currentCandidate.Bid = bid
		currentCandidate.ProcessID = int32(server.port)
	}

	//pass along to neighbour
	neighbour := servers[8080+((*serverId-1)%3)]
	neighbour.Election(context.Background(), currentCandidate)

	//wait for timeout

	//pass further if timeout

	return &proto.Acknowledgement{}, nil
}

func AnnounceResult(currentCandidate *proto.RingLeaderTopDawgG) {
	for _, server := range servers {
		server.Elected(context.Background(), currentCandidate)
	}
}

func connectToServer(port int) {
	conn, _ := grpc.Dial(":"+strconv.Itoa(port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	replica := proto.NewAuctionClient(conn)
	servers[port] = replica

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
