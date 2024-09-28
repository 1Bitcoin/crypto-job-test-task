package main

import (
	trmSqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	rates "testTask/grpc/rates"
	database "testTask/internal/infrastructure/database/postgres"
	"testTask/internal/rpc/get_rates"
	"testTask/internal/usecase/actual_rate_get/repository"
	"testTask/internal/usecase/actual_rate_get/usecase"
)

func main() {
	db, err := database.NewPostgres()
	if err != nil {
		// TODO logger
		log.Fatalf("failed to connect db: %v", err)
	}

	// repository
	rateRepository := repository.New(db, trmSqlx.DefaultCtxGetter)
	rateUsecase := usecase.New(rateRepository)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		// TODO logger
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	rates.RegisterRateServiceServer(s, &get_rates.RateServiceServer{
		Usecase: rateUsecase,
	})

	// TODO logger
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		// TODO logger
		log.Fatalf("failed to serve: %v", err)
	}
}
