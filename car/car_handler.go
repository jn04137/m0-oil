package car

import (
	"net/http"
)

type CarHandler struct {

}

func NewCarHandler(repo *CarRepository) *CarHandler {
	return &CarHandler{}
}

func (h CarHandler) createCar(w http.ResponseWriter, r *http.Request) {

}

func (h CarHandler) updateCar(w http.ResponseWriter, r *http.Request) {

}
