package main

import (
	"context"
	"fmt"
	"os"

	"github.com/elangreza14/go-pg-tx-repository/model"
	"github.com/elangreza14/go-pg-tx-repository/repository"
	"github.com/google/uuid"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load(".env")
	errChecker(err)

	// zap
	// TODO implement this later
	// https: //betterstack.com/community/guides/logging/go/zap/#hiding-sensitive-details-in-your-logs
	logger := zap.NewExample(zap.IncreaseLevel(zap.InfoLevel))
	defer logger.Sync()

	ctx := context.Background()

	// connect postgres
	dbPool, err := DB(ctx, logger)
	errChecker(err)
	defer dbPool.Close()

	userTokenRepository := repository.NewUserTokenRepository(dbPool)

	err = userTokenRepository.CreateUserTX(context.Background(), &model.User{
		ID:       uuid.New(),
		Username: "a",
		Email:    "a",
		Password: []byte("a"),
	}, func(res *model.User) error {
		return nil
	})
	fmt.Println(err)

}

func errChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func DB(ctx context.Context, logger *zap.Logger) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL"),
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	logLevel, err := tracelog.LogLevelFromString(logger.Level().String())
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxzap.NewLogger(logger),
		LogLevel: logLevel,
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
