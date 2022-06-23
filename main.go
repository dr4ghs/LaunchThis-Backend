package main

import (
	"github.com/agsystem/launchthis-be/api"
	ltdb "github.com/agsystem/launchthis-be/db"
)

func main() {
	ltdb.Setup()
	api.Run()
}
