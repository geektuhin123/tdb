package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	pb "github.com/geektuhin123/tdb/whiteboard"
)

type server struct {
	pb.UnimplementedWhiteboardServer
	mu          sync.Mutex
	strokes     []*pb.Stroke
	current     int
	undoHistory []*pb.Stroke
	redoHistory []*pb.Stroke
}

func (s *server) Draw(ctx context.Context, req *pb.DrawRequest) (*pb.DrawResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stroke := &pb.Stroke{
		Id:        int32(len(s.strokes)),
		Color:     req.GetColor(),
		LineWidth: req.GetLineWidth(),
		Points:    req.GetPoints(),
	}

	s.strokes = append(s.strokes, stroke)
	s.current++

	return &pb.DrawResponse{Id: stroke.GetId()}, nil
}

func (s *server) GetWhiteboard(ctx context.Context, req *empty.Empty) (*pb.GetWhiteboardResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &pb.GetWhiteboardResponse{Strokes: s.strokes}, nil
}

func (s *server) Undo(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.current > 0 {
		stroke := s.strokes[s.current-1]
		s.undoHistory = append(s.undoHistory, stroke)
		s.strokes = s.strokes[:s.current-1]
		s.current--
	}

	return &empty.Empty{}, nil
}

func (s *server) Redo(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.undoHistory) > 0 {
		stroke := s.undoHistory[len(s.undoHistory)-1]
		s.redoHistory = append(s.redoHistory, stroke)
		s.strokes = append(s.strokes, stroke)
		s.current++
		s.undoHistory = s.undoHistory[:len(s.undoHistory)-1]
	}

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWhiteboardServer(s, &server{})

	fmt.Println("Server started at port :9000")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
