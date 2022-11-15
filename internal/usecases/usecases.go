package usecases

import (
	"context"
	"github.com/ninja-dark/GoCloudCamp/internal/db"
	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type ServiceLogic interface {
	CreateServ(m enticonfig.Config) error
	GetConfig(ctx context.Context, name string) (*enticonfig.Config, error)
	DeleteConfig(ctx context.Context, name string) error

}

type ServiceLog struct {
	repository db.Repository
}

func (l *ServiceLog) CreateServ(m enticonfig.Config) error {
	
	if err := l.repository.CreateService(m.Service, m.MyData); err != nil {
		return err
	}

	return nil

}

func (l *ServiceLog) GetConfig(ctx context.Context, name string) ( *enticonfig.Config, error) {
	r := l.repository
	return r.GetConfig(ctx, name)

}
func (l *ServiceLog) DeleteConfig(ctx context.Context, name string) error{
	if err := l.repository.DeleteConfig(ctx, name); err != nil {
		return err
	}

	return nil

}

func New(repositpry db.Repository) ServiceLogic {
	return &ServiceLog{repository: repositpry}
}
