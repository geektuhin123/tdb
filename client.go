package main

import (
	"context"
	"log"

	pb "github.com/geektuhin123/tdb/whiteboard"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewWhiteboardClient(conn)

	points := []*pb.Point{
		{X: 10, Y: 10},
		{X: 20, Y: 20},
		{X: 30, Y: 30},
	}

	req := &pb.DrawRequest{
		Color:     "blue",
		LineWidth: 2,
		Points:    points,
	}

	res, err := c.Draw(context.Background(), req)
	if err != nil {
		log.Fatalf("could not draw: %v", err)
	}

	log.Printf("stroke id: %v", res.Id)
}
