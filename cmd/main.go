package main

import(
	"log"

	todoapp "github.com/BogdanAziev2003/ToDoApp"
	"github.com/BogdanAziev2003/ToDoApp/pkg/handler"
)


func main() {
	handlers := new(handler.Handler)
	srv := new(todoapp.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurated while running http server %s", err.Error())
	}
}
