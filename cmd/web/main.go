package main

import (
	"github.com/VATUSA/discord-bot-v3/internal/web"
)

func main() {
	e := web.App()
	e.Logger.Fatal(e.Start(":9002"))
}
