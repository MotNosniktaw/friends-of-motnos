package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/motnosniktaw/nflfantasy/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "TEAMS PORT")
)

type server struct {
	pb.TeamServer
}

func (s *server) Get(ctx context.Context, in *pb.GetTeamRequest) (*pb.GetTeamResponse, error) {
	log.Printf("Received: %v", in.Id)
	return &pb.GetTeamResponse{Id: in.Id, OwnerId: "Yo Mama!"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTeamServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
