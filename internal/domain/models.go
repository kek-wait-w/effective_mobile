package domain

type Car struct {
	ID      int    `json:"id,omitempty"`
	RegNum  string `json:"regNum"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerID int    `json:"ownerId"`
}

type People struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Page struct {
	Cars []Car `json:"cars"`
}

type RegNumsRequest struct {
	RegNums []string `json:"regNums"`
}

type CarInfoResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

type CarsUsecase interface {
	GetCarsPage(filters map[string]string) (page []Page, err error)
	DeleteCar(id int) error
	UpdateCarInfo(vars map[string]string) (car Car, err error)
	PostNewCars(regNums []string) error
}

type CarsRepository interface {
	GetCarsList(filters map[string]string) (cars []Car, err error)
	DeleteCar(id int) error
	UpdateCarInfo(id int, vars map[string]string) (car Car, err error)
	PostNewCars(carsInfo []CarInfoResponse) error
}
