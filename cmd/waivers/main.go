package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/motnosniktaw/nflfantasy/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultTeamId = "potato"
)

var (
	addr   = flag.String("addr", "localhost:50051", "teams server address")
	teamId = flag.String("teanId", defaultTeamId, "team id")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTeamClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &pb.GetTeamRequest{Id: *teamId})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("TeamId: %s", r.GetId())
	log.Printf("OwnerId: %s", r.GetOwnerId())
}
