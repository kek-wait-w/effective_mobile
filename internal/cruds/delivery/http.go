package delivery

import (
	"effect/internal/domain"
	logs "effect/internal/logger"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CarsHandler struct {
	CarsUsecase domain.CarsUsecase
}

func NewCarsHandler(router *mux.Router, u domain.CarsUsecase) {
	handler := &CarsHandler{
		CarsUsecase: u,
	}

	router.HandleFunc("/api/v1/cars/get", handler.GetCarsList).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/cars/delete/{id}", handler.DeleteCar).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/api/v1/cars/update", handler.UpdateCarInfo).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/api/v1/cars/new", handler.PostNewCars).Methods(http.MethodPost, http.MethodOptions)
}

// GetCarsList godoc
//
//	@Summary		Get cars list
//	@Description	Get list of cars by filters.
//	@Tags			cars
//	@Produce		json
//
//	@Success		200		{json}	object{body=object{[]domain.Page}}
//	@Failure		400		{json}	object{err=string}
//
//	@Router			/api/v1/cars/get [get]
func (h *CarsHandler) GetCarsList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	filters := make(map[string]string)

	for key, value := range vars {
		if len(value) > 0 {
			filters[key] = value[0]
		}
	}

	pages, err := h.CarsUsecase.GetCarsPage(filters)

	if err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "GetCarsList", err, err.Error())
		return
	}

	logs.Logger.Debug("http GetCarsList", pages)

	domain.WriteResponse(
		w,
		map[string]interface{}{
			"pages": pages,
		},
		http.StatusOK,
	)

}

// DeleteCar godoc
//
//	@Summary		Delete car
//	@Description	Delete information about car by car id.
//	@Tags			cars
//	@Produce		json
//
//	@Success		200		{json}	object{body=object{"successful removal"}}
//	@Failure		400		{json}	object{err=string}
//
//	@Router			/api/v1/cars/delete/{id} [delete]
func (h *CarsHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "DeleteCar", err, err.Error())
		return
	}

	err = h.CarsUsecase.DeleteCar(id)

	if err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "DeleteCar", err, err.Error())
		return
	}

	domain.WriteResponse(
		w,
		map[string]interface{}{
			"request": "successful removal",
		},
		http.StatusOK)

}

// UpdateCarInfo godoc
//
//	@Summary		Update info
//	@Description	Update info about car by car id and params.
//	@Tags			cars
//	@Produce		json
//
//	@Success		200		{json}	object{body=object{domain.Car}}
//	@Failure		400		{json}	object{err=string}
//
//	@Router			/api/v1/cars/update [patch]
func (h *CarsHandler) UpdateCarInfo(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	params := make(map[string]string)

	for key, value := range vars {
		if len(value) > 0 {
			params[key] = value[0]
		} else {
			domain.WriteError(w, "ERROR: No params", http.StatusBadRequest)
			logs.LogError(logs.Logger, "http", "GetCarsList", errors.New("No params"), "")
			return
		}
	}

	car, err := h.CarsUsecase.UpdateCarInfo(params)

	if err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "GetCarsList", err, err.Error())
		return
	}

	logs.Logger.Debug(car)

	domain.WriteResponse(
		w,
		map[string]interface{}{
			"car": car,
		},
		http.StatusOK,
	)
}

// PostNewCars godoc
//
//	@Summary		Add new cars.
//	@Description	Add new cars with info about owner.
//	@Tags			cars
//	@Produce		json
//
//	@Success		200		{json}	object{body=object{"successful post"}}
//	@Failure		400		{json}	object{err=string}
//
//	@Router			/api/v1/cars/new [post]
func (h *CarsHandler) PostNewCars(w http.ResponseWriter, r *http.Request) {
	var request domain.RegNumsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "PostNewCars", err, err.Error())
		return
	}

	err := h.CarsUsecase.PostNewCars(request.RegNums)

	if err != nil {
		domain.WriteError(w, err.Error(), http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "PostNewCars", err, err.Error())
		return
	}

	domain.WriteResponse(
		w,
		map[string]interface{}{
			"request": "successful post",
		},
		http.StatusOK)
}
