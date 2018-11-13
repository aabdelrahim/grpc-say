package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/aabdelrahim/grpc-say/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// lookup parse and flag
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav sound output")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != err {
		log.Fatalf("could not dial backend %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	request := &pb.SayRequest{Text: flag.Arg(0)}
	response, err := client.Say(context.Background(), request)
	if err != nil {
		log.Fatalf("could not say %s: %v", request.Text, err)
	}
	err = ioutil.WriteFile(*output, response.Audio, 777)
	if err != nil {
		log.Fatalf("could not write to file %s: %v", *output, err)
	}

}
