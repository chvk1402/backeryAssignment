package server

import (
	"backery/router"
)

func Run() error {
	r := router.Init()
	return r.Run(":8080")
}
