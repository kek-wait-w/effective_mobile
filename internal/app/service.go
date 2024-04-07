package app

import (
	http_ "effect/internal/cruds/delivery"
	post "effect/internal/cruds/repository"
	use "effect/internal/cruds/usecase"
	logs "effect/internal/logger"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func Start() {
	password := os.Getenv("POSTGRES_PASSWORD")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Logger.Fatal(logs.Logger, " main ", err, " Failed to connect database")
	}

	router := mux.NewRouter()

	cr := post.NewCarsPostgresqlRepository(db)
	cu := use.NewCarsUsecase(cr)
	http_.NewCarsHandler(router, cu)

	logs.Logger.Info("starting server at :" + os.Getenv("SERVER_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router)
	if err != nil {
		logs.Logger.Fatal(logs.Logger, " main ", err, " Failed to start server")
	}
	logs.Logger.Info("server stopped")
}
