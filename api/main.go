package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"cryptoCurrencies/Database"
	"cryptoCurrencies/Models"
	"cryptoCurrencies/Routes"
	"cryptoCurrencies/Scheduler"
)

var err error

const (
	DBName       = "crypto"
	dbParamsFile = "Config/DatabaseParams.csv"
	dbConfigFile = "Config/DatabaseConfig.csv"
)

func init() {

	//Open DB Connection
	Database.DB, err = gorm.Open("mysql",
		Database.BuildDBConfig(dbConfigFile)+Database.BuildDBParams(dbParamsFile))

	if err != nil {
		log.Fatalf("Error while connecting to DB : %v", err)
	}

	//Create DB if it doesn't exist
	Database.DB.Exec("CREATE DATABASE IF NOT EXISTS " + DBName)
	Database.DB.Exec("Use " + DBName)

	//Create Table with Models.Pairs as the schema
	Database.DB.AutoMigrate(&Models.Pairs{})
}

func Shutdown() {
	// Invoked on Clean Shutdown
	Database.DB.Close()
	Scheduler.StopDBSync()
}

func main() {

	log.Println("Starting Crypto server")

	//Setup router and Start HTTP server
	r := Routes.SetupRoutes()

	srv := &http.Server{
		Addr:    ":8011",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Initiated...")
	Shutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
