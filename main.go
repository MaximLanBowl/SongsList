package main

import (
	_ "SongsList/docs"
	"SongsList/pkg/handlers"
	"SongsList/pkg/repository"

	"os"

	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


var ErrToStart = errors.New("failed to start server")

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed to init config: %s", err)
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error to loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.ConfigToConnect {
		Host: viper.GetString("db.Host"),
		Port: viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname: viper.GetString("db.DBname"),
		SSLmode: viper.GetString("db.SSLmode"),
	})

	repository.DB = db

	if err != nil {
		logrus.Fatalf("Failed to connection: %s", err.Error())
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/songs", handlers.GetSongs) // @Summary Get list of songs
	router.GET("/songs/:id", handlers.GetSongByID) // @Summary Get songs by ID
	router.POST("/songs", handlers.CreateSong) // @Summary Create a new song
	router.PUT("/songs/:id", handlers.UpdateSong) // @Summary Update song by ID
	router.DELETE("/songs/:id", handlers.DeleteSong) // @Summary Delete song by ID
	
	if err := router.Run(":" + viper.GetString("port")); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

