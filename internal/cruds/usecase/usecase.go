package usecase

import (
	"effect/internal/domain"
	logs "effect/internal/logger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type CarsUsecase struct {
	repo domain.CarsRepository
}

func NewCarsUsecase(r domain.CarsRepository) domain.CarsUsecase {
	return &CarsUsecase{
		repo: r,
	}
}

func (u *CarsUsecase) GetCarsPage(filters map[string]string) (result []domain.Page, err error) {
	var pageCount, pageSize int

	for key, value := range filters {
		if key == "page" {
			pageCount, err = strconv.Atoi(value)

			if err != nil {
				logs.LogError(logs.Logger, "usecase", "GetCarsList", err, err.Error())
				return nil, err
			}

			delete(filters, key)
		}
		if key == "pageSize" {
			pageSize, err = strconv.Atoi(value)

			if err != nil {
				logs.LogError(logs.Logger, "usecase", "GetCarsList", err, err.Error())
				return nil, err
			}

			delete(filters, key)
		}
	}
	cars, err := u.repo.GetCarsList(filters)

	if err != nil {
		logs.LogError(logs.Logger, "usecase", "GetCarsList", err, err.Error())
		return nil, err
	}

	if len(cars) == 0 {
		logs.LogError(logs.Logger, "usecase", "GetCarsList", fmt.Errorf("No cars "), "No cars")
		return nil, errors.New("requested Item is not found")
	}

	var pages []domain.Page

	for i := 0; i < pageCount; i++ {
		var page domain.Page
		for j := 0; j < pageSize; j++ {
			if j == len(cars) {
				break
			}
			page.Cars = append(page.Cars, cars[j])
		}
		pages = append(pages, page)
	}

	return pages, nil

}

func (u *CarsUsecase) DeleteCar(id int) error {
	err := u.repo.DeleteCar(id)
	if err != nil {
		logs.LogError(logs.Logger, "usecase", "DeleteCar", err, err.Error())
		return err
	}
	return nil
}

func (u *CarsUsecase) UpdateCarInfo(params map[string]string) (result domain.Car, err error) {
	var id int
	for key, value := range params {
		if key == "id" {
			id, err = strconv.Atoi(value)

			if err != nil {
				logs.LogError(logs.Logger, "usecase", "UpdateCarInfo", err, err.Error())
				return domain.Car{}, err
			}

			delete(params, key)
		}
	}

	car, err := u.repo.UpdateCarInfo(id, params)

	if err != nil {
		logs.LogError(logs.Logger, "usecase", "UpdateCarInfo", err, err.Error())
		return domain.Car{}, err
	}

	return car, err
}

func (u *CarsUsecase) PostNewCars(regNums []string) error {

	var carsInfo []domain.CarInfoResponse

	for _, regNum := range regNums {
		apiUrl := fmt.Sprintf("http://some-api.com/info?regNum=%s", regNum)
		resp, err := http.Get(apiUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var carInfo domain.CarInfoResponse
		if err = json.NewDecoder(resp.Body).Decode(&carInfo); err != nil {
			return err
		}

		carsInfo = append(carsInfo, carInfo)
	}

	err := u.repo.PostNewCars(carsInfo)
	if err != nil {
		logs.LogError(logs.Logger, "usecase", "PostNewCars", err, err.Error())
		return err
	}

	return nil
}
