package car

import (
	"github.com/jmoiron/sqlx"
)

type CarRepository struct {
	db *sqlx.DB 
}

func NewCarRepository(DB *sqlx.DB) *CarRepository {
	return &CarRepository{
		db: DB,
	}
}

func (r CarRepository) CreateCar() {
}

func (r CarRepository) UpdateCar() {
}
