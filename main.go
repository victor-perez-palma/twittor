package main

import (
	"log"

	"github.com/victor-perez-palma/twittor/bd"
	"github.com/victor-perez-palma/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion  a la BD")
		return
	}
	handlers.Manejadores()
}
