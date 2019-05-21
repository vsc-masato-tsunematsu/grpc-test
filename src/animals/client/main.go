package main

import (
  "context"
  "log"
  "os"
  "time"

  "google.golang.org/grpc"
  pb "animals/animals"
)

const (
  address     = "server:50051"
  defaultName = "Mike"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewCatClient(conn)

    name := defaultName
    if len(os.Args) > 1 {
      name = os.Args[1]
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.GetMyCat(ctx, &pb.GetMyCatRequest{Name: name})
    if err != nil {
      log.Fatalf("Error: %v", err)
    }
    log.Printf("Response.Message: %s", r.Message)
}
