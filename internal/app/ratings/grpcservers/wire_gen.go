// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/app/ratings/services"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
)

// Injectors from wire.go:

func CreateRatingsServer(cf string, service services.RatingsService) (*RatingsServer, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	ratingsServer, err := NewRatingsServer(logger, service)
	if err != nil {
		return nil, err
	}
	return ratingsServer, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)