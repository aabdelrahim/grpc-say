package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("could not listen to port %d: %v", *port, err)
	}
	server := grpc.NewServer()
}

// cmd := exec.Command("flite", os.Args[1], "-o", "output.wav")
// cmd.Stderr = os.Stderr
// cmd.Stdout = os.Stdout
// err := cmd.Run()
// if err != nil {
// 	log.Fatal(err)
// }
