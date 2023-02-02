package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/nandohos/go-warehouse-manager/logistics/domain/dto"
)

type ITruckRepository interface {
	GetByRegistration(ctx context.Context, registration string) (*dto.Truck, error)
}

type truckConnection struct {
	connection *sqlx.DB
}

func NewTruckRepository(db *sqlx.DB) ITruckRepository {
	return &truckConnection{
		connection: db,
	}
}

func (t truckConnection) GetByRegistration(ctx context.Context, registration string) (*dto.Truck, error) {
	//TODO implement me
	panic("implement me")
}
