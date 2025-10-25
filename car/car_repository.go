package car

import (
	"github.com/jmoiron/sqlx"
)

type CarRepository struct {
	DBConn *sqlx.DB 
}

func NewCarRepository() *CarRepository {
	return &CarRepository{}
}

func (cr CarRepository) CreateCar() {

}

func (cr CarRepository) UpdateCar() {

}
