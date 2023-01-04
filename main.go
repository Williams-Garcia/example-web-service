package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	router := gin.Default()
	pingpong(router)
	sayHi(router)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func pingpong(router *gin.Engine) {

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}

func sayHi(router *gin.Engine) {
	hi := "Hola %s %s"

	router.POST("/hi", func(ctx *gin.Context) {
		var p Person
		if err := ctx.BindJSON(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Datos de entrada no validos"})
		}

		msg := fmt.Sprintf(hi, p.Name, p.LastName)
		ctx.JSON(http.StatusOK, gin.H{"Data": msg})
	})

}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// /* Definimos la función home */
// func home(w http.ResponseWriter, r *http.Request) {
// 	html := "<html>"
// 	html += "<body>"
// 	html += "<h1>Hola Mundo</h1>"
// 	html += "</body>"
// 	html += "</html>"
// 	w.Write([]byte(html))
// }

// func main() {
// 	jsonMap := map[string]any{
// 		"k1": "val",
// 		"k2": true,
// 	}
// 	fmt.Println(jsonMap)
// 	//marshall casi siempre devolverlo en inline
// 	//el indent lo devuelve identado pero malo
// 	mapAsJson, err := json.MarshalIndent(jsonMap, "", " ")
// 	if err != nil {
// 		fmt.Println()
// 	}
// 	fmt.Println(string(mapAsJson))
// 	/* Definimos la ruta que llamará la función home */
// 	// http.HandleFunc("/", home)
// 	// http.ListenAndServe(":8080", nil)

// }
