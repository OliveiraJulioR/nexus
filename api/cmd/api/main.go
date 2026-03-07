package main

import (
	"github.com/OliveiraJulioR/nexus/api/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
