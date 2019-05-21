package main

import (
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "animals/animals"
)

const (
  port = ":50051"
)

type CatServer struct{}

func (s *CatServer) GetMyCat(ctx context.Context, request *pb.GetMyCatRequest) (*pb.MyCatResponse, error) {
  return &pb.MyCatResponse{Message: request.Name + ": mew!"}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterCatServer(s, &CatServer{})
  s.Serve(lis)
}
