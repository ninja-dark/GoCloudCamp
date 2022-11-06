package db

import (
	"context"

	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type Repository interface{
	InitSchema(ctx context.Context) error
	CreateService(ctx context.Context, name string, data []enticonfig.MyData) ([]enticonfig.MyData, error)
	CreateMessage(ctx context.Context, key string, value string) (*enticonfig.MyData, error)
	GetConfig(ctx context.Context, name string)(*enticonfig.MyData, error)
	DeleteConfig(ctx context.Context, name string) (*enticonfig.MyData, error)
}