package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"review/proto"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type reviewServer struct {
	proto.UnimplementedReviewServer
}

func (rs *reviewServer) ProcessReview(ctx context.Context, review *proto.MyReview) (*proto.MyResponse, error) {

	log.Println("Recieved review:")
	log.Printf("User ID: %d", review.GetUserid())
	log.Printf("Review message: %s", review.GetMessage())
	log.Printf("Rating: %s", review.GetRating().String())

	// Do something complex here

	return &proto.MyResponse{
		Saved: false,
	}, nil
}

func startGrpcServer(ctx context.Context, waitgroup *sync.WaitGroup) error {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	server := grpc.NewServer()
	proto.RegisterReviewServer(server, &reviewServer{})

	// Pre-emptively add 1 to the waitgroup
	waitgroup.Add(1)
	go func() {
		log.Println("Starting gRPC server")
		if err := server.Serve(lis); lis != nil {
			// Decrement waitgroup
			waitgroup.Done()
			log.Fatalf("Failed to serve: %v", err)
		}

	}()

	// Wait for context to be done
	<-ctx.Done()

	// Block until server stops
	server.GracefulStop()

	// Decrement waitgroup
	waitgroup.Done()

	return nil
}

func serve() {
	// Set up channel to recieve signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM, os.Interrupt)

	// Create a "context with cancellation"??
	ctx, cancel := context.WithCancel(context.Background())

	// Create a waitgroup so we can wait for everything to finish
	var shutdownWaitGroup sync.WaitGroup

	// Start server in a goroutine
	go func() {
		if err := startGrpcServer(ctx, &shutdownWaitGroup); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Wait for sigterm from shell
	<-signalChannel

	// Cancel the context so the gRPC server stops
	cancel()

	// Wait for the inner shutdown group to finish
	shutdownWaitGroup.Wait()

	log.Println("Bye")
}

func call() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	// Close connection whenever this function returns
	defer conn.Close()

	client := proto.NewReviewClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel context whenever this function returns
	defer cancel()

	// Create a message - we reference this by pointer in the MyReview struct
	message := "Kangaroo Jack is my favourite film, despite the 4% Rotten Tomatoes rating"

	log.Println("Sending review...")

	client.ProcessReview(ctx,
		&proto.MyReview{
			Userid:    1234,
			Timestamp: uint64(time.Now().Unix()),
			Message:   &message,
			Rating:    proto.Rating_EXCELLENT.Enum(),
		},
	)

}

func main() {

	args := os.Args[1:]

	if args[0] == "send" {
		call()
	} else if args[0] == "recv" {
		serve()
	}
}
