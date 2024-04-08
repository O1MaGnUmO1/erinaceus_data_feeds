package main

import (
	"erinaceus_data_feeds/application"

	"github.com/sirupsen/logrus"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		logrus.Fatalf("failed to load application : reason %v", err)
	}
	app.Run()
}
