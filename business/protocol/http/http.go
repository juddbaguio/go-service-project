package http

import (
	"log"
	"os"
)

type Server interface {
	ListenAndServe() error
}

func Start(container Server) {
	if err := container.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
