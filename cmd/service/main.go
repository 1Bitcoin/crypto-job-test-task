package main

import (
	trmSqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	healthcheckServer "testTask/grpc/healthcheck"
	ratesServer "testTask/grpc/rates"
	"testTask/internal/grpc/getrates"
	"testTask/internal/grpc/healthcheck"
	database "testTask/internal/infrastructure/database/postgres"
	"testTask/internal/infrastructure/env"
	"testTask/internal/usecase/actual_rate_get/repository"
	"testTask/internal/usecase/actual_rate_get/usecase"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	sugar := logger.Sugar()

	err := env.LoadEnv()
	if err != nil {
		sugar.Fatalf("failed load .env file: %v", err)
	}

	db, err := database.NewPostgres()
	if err != nil {
		sugar.Fatalf("failed to connect db: %v", err)
	}

	// repository
	rateRepository := repository.New(db, trmSqlx.DefaultCtxGetter)
	rateUsecase := usecase.New(rateRepository)

	s := grpc.NewServer()
	ratesServer.RegisterRateServiceServer(s, getrates.New(rateUsecase, sugar))
	healthcheckServer.RegisterHealthcheckServiceServer(s, healthcheck.New())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}

	sugar.Infow("Starting gRPC rates on :50051")
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}
