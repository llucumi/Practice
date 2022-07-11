package handler

import (
	"fmt"
	"net/http"
	"practice/internal/domain"
	"practice/internal/users"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type Usuarios struct {
	service users.Service
}

func NewUser(u users.Service) *Usuarios {
	return &Usuarios{
		service: u,
	}
}

func (c *Usuarios) BuscarUsuariosId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, "El id ingresado no existe")
		return
	}
	user, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, valueUser := range user {
		if valueUser.Id == id {
			ctx.JSON(200, gin.H{
				"Usuario encontrado": valueUser,
			})
			return
		}
	}
	ctx.String(400, "No existe ningún usuario con el id ingresado")
}

func (c *Usuarios) FilterUsers(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}
	listUsers, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	route := ctx.Request.URL.Query()

	var filtrado []domain.Usuarios
	for _, u := range listUsers {
		//filtrado por todos los campos
		if route.Get("nombre") == u.Nombre && route.Get("apellido") == u.Apellido && route.Get("email") == u.Email {
			filtrado = append(filtrado, u)
		}
	}
	fmt.Println(filtrado)
	ctx.JSON(http.StatusOK, filtrado)
	//Devuelvo el array filtrado
}

func (c *Usuarios) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *Usuarios) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}
