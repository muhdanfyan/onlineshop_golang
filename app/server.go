package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB

	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on Loading .env.example file")
	}

	appConfig.AppName = "TokoGolang"
	appConfig.AppEnv = "development"
	appConfig.AppPort = "9999"

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
