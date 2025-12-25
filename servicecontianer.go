package main

import (
	"net/http"
	"time"

	"github.com/tanmoyopenroot/upstox-sdk-go/controllers"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/services"
)

// Const ... for http client
const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// Controller ...
type Controller struct {
	*controllers.ConnectController
	*controllers.UserController
	*controllers.OrderController
	*controllers.FeedController
	*controllers.HistoricalController
}

// UpstoxServiceContainer ... Injecting Dependencies
func UpstoxServiceContainer() Controller {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	clientModel := &models.ClientModel{}

	httpService := &services.HTTPService{
		HTTPClient: client,
	}

	connectService := &services.ConnectService{
		HTTPService: httpService,
	}
	connectController := &controllers.ConnectController{
		ConnectService: connectService,
		Client:         clientModel,
	}

	userService := &services.UserService{
		HTTPService: httpService,
	}

	userController := &controllers.UserController{
		UserService: userService,
		Client:      clientModel,
	}

	orderService := &services.OrderService{
		HTTPService: httpService,
	}

	orderController := &controllers.OrderController{
		OrderService: orderService,
		Client:       clientModel,
	}

	feedService := &services.FeedService{
		HTTPService: httpService,
	}

	feedController := &controllers.FeedController{
		FeedService: feedService,
		Client:      clientModel,
	}

	historicalService := &services.HistoricalService{
		HTTPService: httpService,
	}

	historicalController := &controllers.HistoricalController{
		HistoricalService: historicalService,
		Client:            clientModel,
	}

	return Controller{
		connectController,
		userController,
		orderController,
		feedController,
		historicalController,
	}
}
