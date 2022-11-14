package pgrepository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ninja-dark/GoCloudCamp/internal/db"
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

func (r *pgRepository) GetConfig(ctx context.Context, name string)(map[string]string, error){
	rows, _ := r.pool.Query(ctx, ServisesSelect, name)
	//var d map[string]string
	var (
		d map[string]string
		found bool
	)
	for rows.Next(){
		if found{
			return nil, fmt.Errorf("cannot found servises: %s", name)
		}
		err := rows.Scan(name)
			if err != nil{
				return nil, err
			}
		

		found = true
	}
	

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("cannot found servises: %s", name)
	}

	return d, nil
}

func (r *pgRepository) DeleteConfig(ctx context.Context, name string) error{
	rows, _ := r.pool.Query(ctx, DeleteServises, name)
	var (
		found bool
	)
	for rows.Next(){
		if found{
			return  fmt.Errorf("cannot found servises: %s", name)
		}
		if err := rows.Scan(name); err != nil {
			return  err
		}

		found = true
	}

	if err := rows.Err(); err != nil {
		return  err
	}

	if !found {
		return fmt.Errorf("cannot found servises: %s", name)
	}

	return  nil
}

func New(pool *pgxpool.Pool) db.Repository {
	return &pgRepository{pool: pool}
}
