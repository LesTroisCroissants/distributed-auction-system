package main

import (
	"context"
	"distributed-auction-system/proto"
	"errors"
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
	serverPorts    = [3]int{8080, 8081, 8082}
	servers        = make(map[int]proto.AuctionClient) // map port to connections
	serverId       = flag.Int("id", 1, "server id")
	isElecting     = false
	ongoingAuction = false
	leader         = 8080
	bid            int32
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

	AnnounceAuction(30)

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

		_, err := neighbour.Election(context.Background(), currentCandidate)

		if err == nil {
			break
		}
	}
}

func AnnounceResult(currentCandidate *proto.RingLeaderTopDawgG) {
	for _, server := range servers {
		go func(replica proto.AuctionClient) {
			replica.Elected(context.Background(), currentCandidate)
		}(server)
	}
}

func AnnounceNewBid(newBid *proto.ServerBid) {
	for _, server := range servers {
		go func(replica proto.AuctionClient) {
			replica.UpdateBid(context.Background(), newBid)
		}(server)
	}
}

func AnnounceAuction(duration int) {
	for _, server := range servers {
		go func(replica proto.AuctionClient) {
			replica.StartAuction(context.Background(), &proto.AuctionDuration{Duration: int32(duration)})
		}(server)
	}
	go RunAuction(duration)
}

func RunAuction(duration int) {
	ongoingAuction = true
	time.Sleep(time.Duration(duration) * time.Second)
	ongoingAuction = false
}

func (Server) StartAuction(context.Context, *proto.AuctionDuration) (*proto.Empty, error) {
	go RunAuction(30)
	return &proto.Empty{}, nil
}

func (server *Server) Elected(ctx context.Context, electedLeader *proto.RingLeaderTopDawgG) (*proto.Acknowledgement, error) {
	fmt.Println(strconv.Itoa(int(electedLeader.ProcessID)) + " has been elected!")
	leader = int(electedLeader.ProcessID)
	isElecting = false
	return &proto.Acknowledgement{}, nil
}

// API for client
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
		time.Sleep(500 * time.Millisecond)
	}

	return &proto.NewPrimary{
		Port: int32(leader),
	}, nil
}

// API for client
func (server *Server) Bid(ctx context.Context, clientBid *proto.ClientBid) (*proto.Acknowledgement, error) {
	var err error
	if !ongoingAuction {
		err = errors.New("auction is over")
	} else if clientBid.Bid > bid {
		bid = clientBid.Bid
		AnnounceNewBid(&proto.ServerBid{Bid: bid})
		err = nil
	} else {
		err = errors.New("Bid too low")
	}

	return &proto.Acknowledgement{}, err
}

func (server *Server) UpdateBid(ctx context.Context, serverBid *proto.ServerBid) (*proto.Acknowledgement, error) {
	bid = serverBid.Bid
	return &proto.Acknowledgement{}, nil
}

// API for client
func (server *Server) Result(context.Context, *proto.Empty) (*proto.AuctionStatus, error) {
	var status string
	if ongoingAuction {
		status = "highest bid is " + strconv.Itoa(int(bid))
	} else {
		status = "result is " + strconv.Itoa(int(bid))
	}

	return &proto.AuctionStatus{
		Status: status,
	}, nil
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
