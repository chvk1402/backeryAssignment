package server

import (
	"github.com/backery/router"
)

func Run() error {
	r := router.Init()
	return r.Run(":8080")
}
