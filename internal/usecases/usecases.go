package usecases

import (
	"context"
	"log"

	"github.com/ninja-dark/GoCloudCamp/internal/db"
	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type ServiceLogic interface {
	Create(ctc context.Context, name string, c []enticonfig.MyData) *enticonfig.Config
	GetConfig(ctx context.Context, name string) (*enticonfig.MyData, error)
	DeleteConfig(ctx context.Context, name string)(*enticonfig.MyData, error)

}

type ServiceLog struct {
	repository db.Repository
}

func (l *ServiceLog) Create(ctx context.Context, name string, data []enticonfig.MyData) *enticonfig.Config {
	r := l.repository
	d, err := r.CreateService(ctx, name, data)
	if err != nil{
		log.Printf("cannot insert new data servises")
	}

	c := enticonfig.Config{
		Service: name,
		Data: d,
	}
	return &c
}

func (l *ServiceLog) GetConfig(ctx context.Context, name string) (*enticonfig.MyData, error) {
	return l.repository.GetConfig(ctx, name)

}
func (l *ServiceLog) DeleteConfig(ctx context.Context, name string)(*enticonfig.MyData, error){
	return l.repository.DeleteConfig(ctx, name)

}

func New(repositpry db.Repository) ServiceLogic {
	return &ServiceLog{repository: repositpry}
}
