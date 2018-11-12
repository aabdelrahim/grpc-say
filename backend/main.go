package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aabdelrahim/grpc-say/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()

	pb.RegisterTextToSpeechServer(s, server{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatal("could not serve %v", err)
	}
}

type server struct {
}

func (server) Say(ctx context.Context, req *pb.SayRequest) (*pb.SayResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// cmd := exec.Command("flite", os.Args[1], "-o", "output.wav")
// cmd.Stderr = os.Stderr
// cmd.Stdout = os.Stdout
// err := cmd.Run()
// if err != nil {
// 	log.Fatal(err)
// }
