package car

type Car struct {
	NanoId string `json:"nanoId"`
	Make string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
	Vin string `json:"vin"`
}

func NewCar(nanoId string, make string, model string, year int, vin string) Car {
	return Car{
		nanoId,
		make,
		model,
		year,
		vin,
	}
}
