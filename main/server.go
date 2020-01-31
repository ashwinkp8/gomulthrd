package main

import (
	"context"
	"math/rand"
	"net"
	"time"

	pb "github.com/ashwinkp8/gomulthrd/ollo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayOllo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	random := rand.Intn(20)
	time.Sleep(time.Duration(random) * time.Second)

	r := pb.Response{
		FullGreet: "Ollo, " + req.Name,
	}
	return &r, nil
}

func main() {

	l, e := net.Listen("tcp", ":3330")
	if e != nil {
		logrus.Fatalf("could not start. Reason: %v", e)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterSServer(s, &server{})
	s.Serve(l)

}
