package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"grpcProject/pkg/api"
	"grpcProject/pkg/cache"
	"grpcProject/pkg/config"
	"grpcProject/pkg/logger"
	"grpcProject/pkg/repository"
	"grpcProject/pkg/server"
	"log"
	"net"
)

const (
	AddressForListen = "127.0.0.1:8000"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	db, err := repository.NewPostgresDataBase(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	clickhouse, err := repository.NewClickhouse()
	if err != nil {
		return
	}
	defer clickhouse.Close()

	redis, err := cache.NewRedis()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer redis.Close()

	kafka, err := logger.NewKafka()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer kafka.Close()
	grpcServer := grpc.NewServer()
	api.RegisterClientbaseServer(grpcServer, &server.GRPCServer{
		Repository:           repository.NewRepository(db, cfg.DB),
		AdditionalRepository: repository.NewRepository(clickhouse, cfg.AdditionalDB),
		Cache:                cache.NewCache(redis),
		Logger:               logger.NewLogger(kafka),
	})

	listener, err := net.Listen("tcp", AddressForListen)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("listening at %v", AddressForListen)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
		return
	}
}
