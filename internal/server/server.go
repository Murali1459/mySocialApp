package server

import (
	"fmt"
	"mySocialApp/internal/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   database.Db
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}
	Db, err := database.New()
	if err != nil {
		fmt.Println("Error in Creating db")
		return nil
	}
	NewServer.db = *Db

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
