package main

import (
	"github.com/agsystem/launchthis-be/api"
	"github.com/agsystem/launchthis-be/db"
)

func main() {
	db.WaitInit()
	api.WaitInit()
	api.Run()
}
