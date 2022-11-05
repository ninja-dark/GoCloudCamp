package usecases

import (
	"context"

	"github.com/ninja-dark/GoCloudCamp/internal/db"
	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type ServiceLogic interface {
	Create(ctx context.Context, c *enticonfig.Config) (*enticonfig.Config, error)
	//Read(ctx context.Context, id int) (*enticonfig.Config, error)
	//Delete(ctx context.Context, id int) error

}

type ServiceLog struct {
	repository db.Repository
}

func (l *ServiceLog) Create(ctx context.Context, c *enticonfig.Config) (*enticonfig.Config, error) {
	return l.repository.CreateService(ctx, c)
}

func New(repositpry db.Repository) ServiceLogic {
	return &ServiceLog{repository: repositpry}
}
