package server

import (
	"context"
	"fmt"
	"grpcProject/pkg/api"
	"grpcProject/pkg/cache"
	"grpcProject/pkg/logger"
	"grpcProject/pkg/repository"
)

// GRPCServer ...
type GRPCServer struct {
	Repository           *repository.Repository
	AdditionalRepository *repository.Repository
	Cache                *cache.ClientCache
	Logger               *logger.Logger
	api.UnimplementedClientbaseServer
}

// Add : add new user
func (s *GRPCServer) Add(ctx context.Context, req *api.AddClientRequest) (*api.AddClientResponse, error) {

	err := s.Repository.Add(repository.ClientEntity{
		Name: req.Client.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("server can't add client to database: %v", err)
	}
	err = s.AdditionalRepository.Add(repository.ClientEntity{
		Name: req.Client.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("server can't add client to additional database: %v", err)
	}
	s.Cache.Add(cache.ClientEntity{
		Value: req.Client.Name,
	})

	return &api.AddClientResponse{}, nil
}

// Delete : delete a user with a specific name
func (s *GRPCServer) Delete(ctx context.Context, req *api.DeleteClientRequest) (*api.DeleteClientResponse, error) {
	err := s.Repository.Delete(repository.ClientEntity{
		Name: req.Client.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("server can't delete client from repository: %v", err)
	}

	s.Cache.Delete(cache.ClientEntity{
		Value: req.Client.Name,
	})

	return &api.DeleteClientResponse{}, nil
}

// GetClients : get all clients
func (s *GRPCServer) GetClients(context.Context, *api.GetClientsRequest) (*api.GetClientsResponse, error) {
	var clients []*api.Client
	rows, err := s.Repository.GetClients()
	if err != nil {
		return nil, fmt.Errorf("server can't get clients from repository: %v", err)
	}
	for _, row := range rows {
		clients = append(clients, &api.Client{
			Name: row.Name,
		})
	}
	return &api.GetClientsResponse{Clients: clients}, nil
}

// GetClientsFromRedis : get all clients in a minute from redis
func (s *GRPCServer) GetClientsFromRedis(context.Context, *api.GetClientsRequest) (*api.GetClientsResponse, error) {

	result, err := s.Cache.GetClients()
	if err != nil {
		return nil, fmt.Errorf("server can't get cash result: %v", err)
	}

	var clients []*api.Client
	for _, r := range result {
		clients = append(clients, &api.Client{Name: r.Value})
	}

	return &api.GetClientsResponse{Clients: clients}, nil
}
