package main

import (
	"log"
	"net"

	"github.com/sm888sm/task-management-api/authentication-service/config"
	"github.com/sm888sm/task-management-api/authentication-service/internal/controllers"
	"github.com/sm888sm/task-management-api/authentication-service/internal/proto/auth"
	"github.com/sm888sm/task-management-api/authentication-service/internal/repositories"
	"github.com/sm888sm/task-management-api/authentication-service/internal/services"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
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

	// create repository
	// TODO Connect DB
	userRepository := repositories.NewUserRepository()

	// Create an instance of the authentication & token service
	authService := services.NewAuthService(*userRepository)
	tokenService := services.NewTokenService(config)

	// Create an instance of the authentication controller
	authController := controllers.NewAuthController(authService, tokenService)

	// Register the authentication service with the gRPC server
	auth.RegisterAuthServiceServer(server, authController)

	log.Printf("Authentication service is running on port %s", config.Port)

	// Start the gRPC server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
