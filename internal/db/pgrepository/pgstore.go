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
			id INT UNIQUE GENERATED ALWAYS AS IDENTITY,
			name varchar(20) NOT NULL,
			key varchar(20) NOT NULL,
      		value varchar(20) NOT NULL
		);

	`
		ServisesSelect = `SELECT name, key, value FROM servises WHERE name = $1` 
		DeleteServises = `DELETE * FROM servises, message WHERE name = $1`
		AddServises = `
		INSERT INTO servises (
			name, key, value
		) VALUES(
			$1, -- key
			$2, -- value
			$3 -- name
		)
	`
)
var (
	ctx = context.Background()
)

type pgRepository struct{
	pool *pgxpool.Pool
}

func (r *pgRepository) InitSchema(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, DDL)
	return err
}

func (r *pgRepository) CreateService(name string, m map[string]string)  error{
	for k, v:= range m{
		_, err := r.pool.Exec(
			ctx,
			AddServises,
			name,
			k,
			v )
		if err != nil {
				return err
			}
	
	}
	return nil
}

func (r *pgRepository) GetConfig(ctx context.Context, name string)(*enticonfig.Config, error){

	dbs:= &enticonfig.Config{}
	rows, err := r.pool.Query(ctx, ServisesSelect, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&dbs.MyData,
		); err != nil {
			return nil, err
		}
	}

	return &enticonfig.Config{
		MyData: dbs.MyData,
	}, nil
}

func (r *pgRepository) DeleteConfig(ctx context.Context, name string) error{
	_, err := r.pool.Exec(ctx, DeleteServises, name)
	
	return err
}

func New(pool *pgxpool.Pool) db.Repository {
	return &pgRepository{pool: pool}
}
