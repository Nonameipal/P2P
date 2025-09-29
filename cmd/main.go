package main

import (
	"github.com/Nonameipal/P2P/internal/configs"
	"github.com/Nonameipal/P2P/internal/controller"
	"github.com/Nonameipal/P2P/internal/db"
	"github.com/Nonameipal/P2P/internal/repository"
	"github.com/Nonameipal/P2P/internal/service"
	"github.com/rs/zerolog"
	"os"
)

// @title P2P marketplace service
// @contact.name API P2P marketplace
// @contact.url https://test.com/
// @contact.email firdavs022006@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("Starting up - Start")
	if err := configs.ReadSettings(); err != nil {
		logger.Error().Err(err).Msg("Error reading settings" + err.Error())
		return
	}

	dbConn, err := db.InitConnection()
	if err != nil {
		logger.Error().Err(err).Msg("Error during database connection initialization: " + err.Error())
		return
	}

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	if err = ctrl.InitRoutes(); err != nil {
		logger.Error().Err(err).Msg("Error during http-service initialization: " + err.Error())
		return
	}

	if err = db.CloseConnection(dbConn); err != nil {
		logger.Error().Err(err).Msg("Error during database connection close: " + err.Error())
		return
	}

	logger.Info().Msg("Starting up - End")
}
