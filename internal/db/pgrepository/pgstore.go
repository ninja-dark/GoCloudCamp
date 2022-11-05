package pgrepository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ninja-dark/GoCloudCamp/internal/db"
	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
)


const(
	DDL = `
		CREATE TABLE IF NOT EXISTS servises 
		( 
			id integer NOT NULL,
			name varchar(20) NOT NULL, 
			version integer NOT NULL,
			Key1 varchar(20) NOT NULL,
			Key2 varchar(20) Not NULL
		);
	`
		ServisesSelect = `SELECT name, key1, key2 FROM servises WHERE name = $1` 
		DeleteServises = `SELECT * FROM servises WHERE name = $1`
		AddServises = `
		INSERT INTO servises (
			id, name, key1, key2
		) VALUES(
			$1, -- id
			$2, -- name
			$3, -- key1
			$4, -- key2
		) `

)

type pgRepository struct{
	pool *pgxpool.Pool
}

func (r *pgRepository) InitSchema(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, DDL)
	return err
}

func (r *pgRepository) CreateService(ctx context.Context, s *enticonfig.Config) (*enticonfig.Config, error) {

	newService := *s
	
	_, err := r.pool.Exec(
		ctx,
		AddServises,
		newService.Id,
		newService.Service,
		newService.Key1,
		newService.Key2)

	return &newService, err
}

func New(pool *pgxpool.Pool) db.Repository {
	return &pgRepository{pool: pool}
}
