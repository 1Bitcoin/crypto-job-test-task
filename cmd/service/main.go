package main

import (
	trmSqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	healthcheckServer "testTask/grpc/healthcheck"
	ratesServer "testTask/grpc/rates"
	"testTask/internal/grpc/getrates"
	"testTask/internal/grpc/healthcheck"
	database "testTask/internal/infrastructure/database/postgres"
	"testTask/internal/infrastructure/env"
	"testTask/internal/usecase/actual_rate_get/repository"
	"testTask/internal/usecase/actual_rate_get/usecase"
	"time"
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

	// usecase
	rateTimeout := 3 * time.Second
	rateUsecase := usecase.New(
		rateRepository,
		&http.Client{
			Timeout: rateTimeout,
		},
	)

	s := grpc.NewServer()
	ratesServer.RegisterRateServiceServer(s, getrates.New(rateUsecase, sugar))
	healthcheckServer.RegisterHealthcheckServiceServer(s, healthcheck.New())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Wait group to wait for cleanup
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// Wait for an interrupt signal
		<-stop
		sugar.Info("Shutting down server...")

		// Gracefully stop the server
		s.GracefulStop()
		wg.Done()
	}()

	sugar.Infow("Starting gRPC rates on :50051")
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}

	// Wait for the cleanup goroutine to finish
	wg.Wait()
	sugar.Info("Server shut down gracefully")
}
