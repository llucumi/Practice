package main

import (
	"practice/internal/users"
	"practice/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	repoUser := users.NewRepository()
	serviceUser := users.NewService(repoUser)
	usuario := handler.NewUser(serviceUser)

	r := gin.Default()
	us := r.Group("/usuarios")
	us.POST("/", usuario.Store())
	us.GET("/filtrousuario", usuario.FilterUsers)
	us.GET("/:id", usuario.BuscarUsuariosId)
	us.GET("/", usuario.GetAll())
	r.Run()
}
