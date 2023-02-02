package services

import (
	"context"
	"github.com/nandohos/go-warehouse-manager/logistics/domain/dto"
	"github.com/nandohos/go-warehouse-manager/logistics/infrastructure/repositories"
)

type ITruckService interface {
	Get(ctx context.Context, registration string) (*dto.Truck, error)
}

type truckService struct {
	TruckRepository repositories.ITruckRepository
}

func NewTruckService(truckRepo repositories.ITruckRepository) ITruckService {
	return &truckService{
		TruckRepository: truckRepo,
	}
}

func (t truckService) Get(ctx context.Context, registration string) (*dto.Truck, error) {
	//TODO implement me
	panic("implement me")
}
