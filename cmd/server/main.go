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

	clickhouse, err := repository.NewClickhouse(cfg)
	if err != nil {
		return
	}
	defer clickhouse.Close()

	redis, err := cache.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer redis.Close()

	kafka, err := logger.NewKafka(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer kafka.Close()
	grpcServer := grpc.NewServer()
	api.RegisterClientbaseServer(grpcServer, &server.GRPCServer{
		Repository:           repository.NewRepository(db, cfg.DB.DriverName),
		AdditionalRepository: repository.NewRepository(clickhouse, cfg.AdditionalDB.DriverName),
		Cache:                cache.NewCache(redis),
		Logger:               logger.NewLogger(kafka),
	})

	listener, err := net.Listen("tcp", cfg.App.Addressforlisten)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("listening at %v", cfg.App.Addressforlisten)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
		return
	}
}
