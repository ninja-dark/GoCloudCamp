package db

import (
	"context"

	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type Repository interface{
	InitSchema(ctx context.Context) error
	CreateService(name string, m map[string]string) error
	GetConfig(ctx context.Context, name string)( *enticonfig.Config, error)
	DeleteConfig(ctx context.Context, name string) error
}