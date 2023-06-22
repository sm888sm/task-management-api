package main

import (
	"log"
	"net"

	"github.com/sm888sm/task-management-api/authentication-service/internal/controllers"
	"github.com/sm888sm/task-management-api/authentication-service/internal/proto/auth"
	"github.com/sm888sm/task-management-api/authentication-service/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Create an instance of the authentication controller
	authController := controllers.NewAuthController()

	// Register the authentication service with the gRPC server
	auth.RegisterAuthServiceServer(server, authController)

	log.Printf("Authentication service is running on port %s", config.Port)

	// Start the gRPC server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
