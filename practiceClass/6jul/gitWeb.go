package main

//usa go routines para las consultas
import "github.com/gin-gonic/gin"

func main() {
	// Crea un router con gin
	//logger y recover para que no se cierre (2 manejadores)
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	//para crear un router, indicamos que el endpoint
	//que vamos a crear es de tipo GET, bajo hello world,
	//para manejar usamos la función que retorna un 200 y un (TERCER MANEJADOR)

	//git.H es lo mismo 	ue map[string] interface{}
	//el Contexto tiene un montón de métodos para las response
	router.GET("/hello-world", func(c *gin.Context) {
		//respuesta tipo JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
		//respuesta tipo string
		//c.String(200, "Hola mundo")
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}

//nos llega un request, nos redirecciona a través de un router, nos llega el middleware, nos para el manejador handler para indicar a donde va y puede haber un middleware para el response
