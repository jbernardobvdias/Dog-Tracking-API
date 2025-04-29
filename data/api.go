package data

import (
	"dog-tracking/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupHttp(port string) {
	log.Println("Setting up HTTP API...")

	http.HandleFunc("/dogs", HttpHanldeDogs)
	http.HandleFunc("/dogs/", HttpHandleDogWithId)

	http.ListenAndServe(":"+port, nil)

	log.Println("HTTP API set up on port " + port)
}

func SetupGin(port string) {
	log.Println("Setting up Gin API...")

	router := gin.Default()

	router.GET("/dogs", GinGetDogs)
	router.POST("/dogs", GinAddDogs)
	router.PUT("/dogs/:id", GinUpdateDogs)
	router.DELETE("/dogs/:id", GinDeleteDogs)

	router.Run(":" + port)

	log.Println("Gin API set up on port " + port)
}

// HTTP Handlers

func HttpHanldeDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		dogs := GetDogs()
		err := json.NewEncoder(w).Encode(dogs)
		if err != nil {
			http.Error(w, "Failed to encode dogs data", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var dog models.Dog

		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &dog)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		AddDogs(dog.Name, dog.Race, dog.Age)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HttpHandleDogWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/dogs/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var dog models.Dog
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &dog)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		UpdateDogs(id, dog.Name, dog.Race, dog.Age)
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		DeleteDogs(id)
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Gin Handlers

func GinGetDogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetDogs())
}

func GinAddDogs(c *gin.Context) {
	var dog models.Dog
	if err := c.BindJSON(&dog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	AddDogs(dog.Name, dog.Race, dog.Age)
	c.Status(http.StatusCreated)
}

func GinUpdateDogs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var dog models.Dog
	if err := c.BindJSON(&dog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	UpdateDogs(id, dog.Name, dog.Race, dog.Age)
	c.Status(http.StatusOK)
}

func GinDeleteDogs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	DeleteDogs(id)
	c.Status(http.StatusOK)
}
