package db

import (
	"context"

	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)

type Repository interface{
	InitSchema(ctx context.Context) error
	CreateService(ctx context.Context, s *enticonfig.Config) (*enticonfig.Config, error)
}