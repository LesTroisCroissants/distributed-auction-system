// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: proto/proto.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auction_UpdateBid_FullMethodName      = "/auction.Auction/UpdateBid"
	Auction_Election_FullMethodName       = "/auction.Auction/Election"
	Auction_Elected_FullMethodName        = "/auction.Auction/Elected"
	Auction_StartAuction_FullMethodName   = "/auction.Auction/StartAuction"
	Auction_Bid_FullMethodName            = "/auction.Auction/Bid"
	Auction_WhoIsNewLeader_FullMethodName = "/auction.Auction/WhoIsNewLeader"
	Auction_Result_FullMethodName         = "/auction.Auction/Result"
)

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	// Internal RPC
	UpdateBid(ctx context.Context, in *ServerBid, opts ...grpc.CallOption) (*Acknowledgement, error)
	Election(ctx context.Context, in *RingLeaderTopDawgG, opts ...grpc.CallOption) (*Acknowledgement, error)
	Elected(ctx context.Context, in *RingLeaderTopDawgG, opts ...grpc.CallOption) (*Acknowledgement, error)
	StartAuction(ctx context.Context, in *AuctionDuration, opts ...grpc.CallOption) (*Empty, error)
	// Client RPC
	Bid(ctx context.Context, in *ClientBid, opts ...grpc.CallOption) (*Acknowledgement, error)
	WhoIsNewLeader(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewPrimary, error)
	// - client ask occasionally about the result
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuctionStatus, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) UpdateBid(ctx context.Context, in *ServerBid, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, Auction_UpdateBid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Election(ctx context.Context, in *RingLeaderTopDawgG, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, Auction_Election_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Elected(ctx context.Context, in *RingLeaderTopDawgG, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, Auction_Elected_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) StartAuction(ctx context.Context, in *AuctionDuration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Auction_StartAuction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Bid(ctx context.Context, in *ClientBid, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, Auction_Bid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) WhoIsNewLeader(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewPrimary, error) {
	out := new(NewPrimary)
	err := c.cc.Invoke(ctx, Auction_WhoIsNewLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuctionStatus, error) {
	out := new(AuctionStatus)
	err := c.cc.Invoke(ctx, Auction_Result_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility
type AuctionServer interface {
	// Internal RPC
	UpdateBid(context.Context, *ServerBid) (*Acknowledgement, error)
	Election(context.Context, *RingLeaderTopDawgG) (*Acknowledgement, error)
	Elected(context.Context, *RingLeaderTopDawgG) (*Acknowledgement, error)
	StartAuction(context.Context, *AuctionDuration) (*Empty, error)
	// Client RPC
	Bid(context.Context, *ClientBid) (*Acknowledgement, error)
	WhoIsNewLeader(context.Context, *Empty) (*NewPrimary, error)
	// - client ask occasionally about the result
	Result(context.Context, *Empty) (*AuctionStatus, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (UnimplementedAuctionServer) UpdateBid(context.Context, *ServerBid) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBid not implemented")
}
func (UnimplementedAuctionServer) Election(context.Context, *RingLeaderTopDawgG) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedAuctionServer) Elected(context.Context, *RingLeaderTopDawgG) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Elected not implemented")
}
func (UnimplementedAuctionServer) StartAuction(context.Context, *AuctionDuration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAuction not implemented")
}
func (UnimplementedAuctionServer) Bid(context.Context, *ClientBid) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServer) WhoIsNewLeader(context.Context, *Empty) (*NewPrimary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoIsNewLeader not implemented")
}
func (UnimplementedAuctionServer) Result(context.Context, *Empty) (*AuctionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_UpdateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).UpdateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_UpdateBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).UpdateBid(ctx, req.(*ServerBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RingLeaderTopDawgG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Election_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Election(ctx, req.(*RingLeaderTopDawgG))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Elected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RingLeaderTopDawgG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Elected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Elected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Elected(ctx, req.(*RingLeaderTopDawgG))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_StartAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionDuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).StartAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_StartAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).StartAuction(ctx, req.(*AuctionDuration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Bid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Bid(ctx, req.(*ClientBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_WhoIsNewLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).WhoIsNewLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_WhoIsNewLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).WhoIsNewLeader(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Result_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auction.Auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBid",
			Handler:    _Auction_UpdateBid_Handler,
		},
		{
			MethodName: "Election",
			Handler:    _Auction_Election_Handler,
		},
		{
			MethodName: "Elected",
			Handler:    _Auction_Elected_Handler,
		},
		{
			MethodName: "StartAuction",
			Handler:    _Auction_StartAuction_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _Auction_Bid_Handler,
		},
		{
			MethodName: "WhoIsNewLeader",
			Handler:    _Auction_WhoIsNewLeader_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _Auction_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}
