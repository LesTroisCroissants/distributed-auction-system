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
	serverId    = flag.Int("id", 1, "server id")
	timeout     = 2 * time.Second //random assumption
	isElecting  = false
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
	isElecting = true

	go func() {
		//check for own identifier
		if currentCandidate.ProcessID == int32(server.port) {
			AnnounceResult(currentCandidate)
			leader = server.port // set yourself as leader
			return
		}

		//compare with own bid and identifier
		if currentCandidate.Bid < bid || (currentCandidate.Bid == bid && currentCandidate.ProcessID < int32(server.port)) {
			currentCandidate.Bid = bid
			currentCandidate.ProcessID = int32(server.port)
		}

		PassElection(ctx, currentCandidate)
	}()

	return &proto.Acknowledgement{}, nil
}

func PassElection(ctx context.Context, currentCandidate *proto.RingLeaderTopDawgG) {
	for i := 0; i < 2; i++ { //try to contact 2 subsequent nodes in ring. Could be modified to loop over all nodes
		neighbour := servers[8080+((*serverId+i)%3)]
		fmt.Println("neighbour:", 8080+((*serverId+i)%3))

		deadlineContext, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
		defer cancel()
		neighbour.Election(deadlineContext, currentCandidate)

		// Check for timeout reached
		if deadlineContext.Err() == nil {
			break
		}
		// pass to next neighbour if timeout
	}
}

func AnnounceResult(currentCandidate *proto.RingLeaderTopDawgG) {
	for _, server := range servers {
		server.Elected(context.Background(), currentCandidate)
	}
}

func (server *Server) Elected(ctx context.Context, electedLeader *proto.RingLeaderTopDawgG) (*proto.Acknowledgement, error) {
	fmt.Println(strconv.Itoa(int(electedLeader.ProcessID)) + " has been elected!")
	leader = int(electedLeader.ProcessID)
	isElecting = false
	return &proto.Acknowledgement{}, nil
}

func (server *Server) WhoIsNewLeader(ctx context.Context, void *proto.Empty) (*proto.NewPrimary, error) {
	if !isElecting {
		isElecting = true
		PassElection(ctx, &proto.RingLeaderTopDawgG{
			ProcessID: int32(server.port),
			Bid:       bid,
		})
	}

	// wait for election to finish
	for isElecting {
		time.Sleep(5 * time.Second)
	}

	return &proto.NewPrimary{
		Port: int32(leader),
	}, nil
}

func (server *Server) Bid(context.Context, *proto.ClientBid) (*proto.Acknowledgement, error) {
	fmt.Println("Bid has been called!")
	return &proto.Acknowledgement{}, nil
}

func (server *Server) UpdateBid(context.Context, *proto.ServerBid) (*proto.Acknowledgement, error) {
	return &proto.Acknowledgement{}, nil
}

func (server *Server) Result(context.Context, *proto.Empty) (*proto.AuctionStatus, error) {
	return nil, nil // todo: needs to return auction status
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
