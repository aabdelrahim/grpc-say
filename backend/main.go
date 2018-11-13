package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "github.com/aabdelrahim/grpc-say/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	log.Printf("listening on port %d \n", *port)
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
	f, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatalf("could not create temp file: %v", err)
	}
	err = f.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close temp file %s: %v", f.Name(), err)
	}
	cmd := exec.Command("flite", req.Text, "-o", f.Name())

	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("flite command failed: %s : %v", data, err)
	}
	data, err = ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatalf("could not read file %s: %v", f.Name(), err)
	}
	return &pb.SayResponse{Audio: data}, nil
}
