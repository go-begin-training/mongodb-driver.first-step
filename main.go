package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-begin-training/mongodb-driver.first-step/mongodb"

	"github.com/gin-gonic/gin"
)

const (
	addr string = "0.0.0.0:8080"

	mongoURI string = "mongodb://127.0.0.1:27017"
	dbName   string = "first-step"
)

func main() {
	/*
		Khởi tạo kết nối với mongodb
	*/
	if err := mongodb.NewConnection(mongoURI, dbName, 30*time.Second); err != nil {
		log.Fatal(err)
	}

	route := gin.Default()

	path := route.Group("/data")
	{
		path.Handle(http.MethodPost, "/set", setInfo)
		path.Handle(http.MethodPost, "/update", updateInfo)
		path.Handle(http.MethodGet, "/get", getInfo)
	}

	route.Run(addr)
}
