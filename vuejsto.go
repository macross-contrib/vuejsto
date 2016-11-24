package main

import (
	//"github.com/macross-contrib/statio"
	"github.com/insionng/macross"
	"github.com/insionng/macross/logger"
	"github.com/insionng/macross/pongor"
	"github.com/insionng/macross/static"
	"github.com/macross-contrib/vuejsto/handlers"
)

func main() {
	m := macross.New()
	m.Use(logger.Logger())
	m.Use(static.Static("public"))
	//m.Use(statio.Serve("/", statio.LocalFile("public", false)))
	m.SetRenderer(pongor.Renderor())

	m.File("/favicon.ico", "public/favicon.ico")
	m.Get("/", handlers.GetMain)
	m.Get("/tasks", handlers.GetTasks)
	m.Post("/task", handlers.PostTask).Put(handlers.PutTask)
	m.Delete("/task/<id>", handlers.DeleteTask)
	//m.Static("/", "public")

	m.Run(":7000")
}
