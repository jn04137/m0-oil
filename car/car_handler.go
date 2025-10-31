package car

import (
	"log"
	"net/http"
	"encoding/json"
)

type CarHandler struct {
	r *CarRepository
}

func NewCarHandler(repo *CarRepository) *CarHandler {
	return &CarHandler{
		r: repo,
	}
}

func (h *CarHandler) testCar(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["something"] = "someValue"

	json.NewEncoder(w).Encode(data);
}

func (h *CarHandler) getMyCars(w http.ResponseWriter, r *http.Request) {
	cars := []Car{}
	car := NewCar("902jal;skdj3", "Subaru", "Forester", 2025, "JPY02931JDS")
	cars = append(cars, car)

	car = NewCar(";ojnfvi2", "Toyota", "Corolla", 2025, "JPY02931JDS")
	cars = append(cars, car)

	err := json.NewEncoder(w).Encode(cars)
	if err != nil {
		log.Printf("Failed encoding payload: %v", err)
	}
}

func (h *CarHandler) getMyCar(w http.ResponseWriter, r *http.Request) {
	carId := r.PathValue("carId")
	log.Printf("This is the car: %v", carId)

	car := NewCar("902jal;skdj3", "Subaru", "Forester", 2025, "JPY02931JDS")
	json.NewEncoder(w).Encode(car)
}

func (h CarHandler) createCar(w http.ResponseWriter, r *http.Request) {
	h.r.CreateCar()
}

func (h *CarHandler) updateCar(w http.ResponseWriter, r *http.Request) {

}

func NewCarRouter(repo *CarRepository) *http.ServeMux {
	h := CarHandler{r: repo}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", h.testCar)
	mux.HandleFunc("GET /garage/mycars", h.getMyCars)
	mux.HandleFunc("GET /garage/mycars/{carId}", h.getMyCar)
	mux.HandleFunc("POST /update", h.updateCar)
	mux.HandleFunc("POST /create", h.createCar)

	return mux
}
