package repository

import (
	"effect/internal/domain"
	logs "effect/internal/logger"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type postgresqlRepository struct {
	db *gorm.DB
}

func NewCarsPostgresqlRepository(db *gorm.DB) domain.CarsRepository {
	return &postgresqlRepository{
		db: db,
	}
}

func (r *postgresqlRepository) GetCarsList(filters map[string]string) (cars []domain.Car, err error) {

	query := r.db.Model(&domain.Car{})

	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if len(filters) == 0 {
		query = query.Find(&cars)
	} else {
		query.Find(&cars)
	}

	return cars, nil
}

func (r *postgresqlRepository) DeleteCar(id int) error {
	car := domain.Car{ID: id}

	result := r.db.Delete(&car)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		logs.LogError(logs.Logger, "repo", "DeleteCar", errors.New("car not found"), "")
		return errors.New("car not found")
	}
	return nil
}

func (r *postgresqlRepository) UpdateCarInfo(id int, vars map[string]string) (domain.Car, error) {
	car := domain.Car{ID: id}

	for key, value := range vars {
		switch key {
		case "mark":
			car.Mark = value
		case "model":
			car.Model = value
		case "regNum":
			car.RegNum = value
		case "year":
			year, err := strconv.Atoi(value)
			if err != nil {
				logs.LogError(logs.Logger, "repo", "UpdateCarInfo", err, err.Error())
				return domain.Car{}, err
			}
			car.Year = year
		}
	}

	err := r.db.Model(&domain.Car{}).Where("id = ?", id).Updates(&car).Error
	if err != nil {
		logs.LogError(logs.Logger, "repo", "UpdateCarInfo", err, err.Error())
		return domain.Car{}, err
	}

	return car, nil
}

func (r *postgresqlRepository) PostNewCars(carsInfo []domain.CarInfoResponse) error {
	for _, carInfo := range carsInfo {
		owner := domain.People{
			Name:       carInfo.Owner.Name,
			Surname:    carInfo.Owner.Surname,
			Patronymic: carInfo.Owner.Patronymic,
		}
		err := r.db.Create(&owner).Error
		if err != nil {
			return err
		}

		car := domain.Car{
			RegNum:  carInfo.RegNum,
			Mark:    carInfo.Mark,
			Model:   carInfo.Model,
			Year:    carInfo.Year,
			OwnerID: owner.ID,
		}
		err = r.db.Create(&car).Error
		if err != nil {
			return err
		}
	}

	return nil
}
