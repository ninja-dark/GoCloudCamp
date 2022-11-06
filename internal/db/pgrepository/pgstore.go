package pgrepository

import (
	"context"
	"fmt"

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
			version integer NOT NULL
		);


		CREATE TABLE IF NOT EXISTS message 
    	(
      		id INT UNIQUE GENERATED ALWAYS AS IDENTITY,
      		Key varchar(20) NOT NULL,
      		Value varchar(20) NOT NULL, 
      		config_id integer NOT NULL,
			FOREIGN KEY (config_id) REFERENCES servises (id)
			  DEFERRABLE INITIALLY DEFERRED
    );
	`
		ServisesSelect = `SELECT name, key1, key2 FROM servises WHERE name = $1` 
		DeleteServises = `DELETE * FROM servises, message WHERE name = $1`
		AddMessage = `
		INSERT INTO message (
			id, key, value, config_id
		) VALUES(
			$1, -- id
			$2, -- key
			$3, -- value
			$4  -- config_id
		)
	`
		AddServises = `
		INSERT INTO servises(
			id, name, version	
		) VALUES (
			$1, -- id
			$2,  -- name
			$3 	-- version
		)
	`

)

type pgRepository struct{
	pool *pgxpool.Pool
}

func (r *pgRepository) InitSchema(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, DDL)
	return err
}

func (r *pgRepository) CreateService(ctx context.Context, name string, data []enticonfig.MyData) ([]enticonfig.MyData, error) {
	v := 0
	
	newService := data
	_, err := r.pool.Exec(
		ctx,
		AddServises,
		name, v)
	for v := range data{
		r.CreateMessage(ctx, data[v].Key, data[v].Value )
	}

	return newService, err
}

func (r *pgRepository)CreateMessage(ctx context.Context, key string, value string) (*enticonfig.MyData, error){

	_, err := r.pool.Exec(
		ctx, 
		AddMessage,
		key,
		value)
	return &enticonfig.MyData{Key: key, Value: value}, err
}

func (r *pgRepository) GetConfig(ctx context.Context, name string)(*enticonfig.MyData, error){
	rows, _ := r.pool.Query(ctx, ServisesSelect, name)
	var (
		s enticonfig.MyData
		found bool
	)
	for rows.Next(){
		if found{
			return nil, fmt.Errorf("cannot found servises: %s", name)
		}
		if err := rows.Scan(name); err != nil {
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

	return &s, nil
}

func (r *pgRepository) DeleteConfig(ctx context.Context, name string)(*enticonfig.MyData, error){
	rows, _ := r.pool.Query(ctx, DeleteServises, name)
	var (
		s enticonfig.MyData
		found bool
	)
	for rows.Next(){
		if found{
			return nil, fmt.Errorf("cannot found servises: %s", name)
		}
		if err := rows.Scan(name); err != nil {
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

	return &s, nil
}

func New(pool *pgxpool.Pool) db.Repository {
	return &pgRepository{pool: pool}
}
