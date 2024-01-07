package main

import (
	"github.com/fvbock/endless"
	"github.com/lixiao189/msgr-bot/app/router"
)

func main() {
	router := router.InitRouter()
	endless.ListenAndServe(":8080", router)
}
