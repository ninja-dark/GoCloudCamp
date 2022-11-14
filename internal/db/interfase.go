package db

import (
	"context"
)

type Repository interface{
	InitSchema(ctx context.Context) error
	CreateService(name string, m map[string]string) error
	GetConfig(ctx context.Context, name string)( map[string]string, error)
	DeleteConfig(ctx context.Context, name string) error
}