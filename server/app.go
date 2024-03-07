package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	biometricHttp "pthw.com/asymmetric-for-biometric/delivery/http/biometric"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
	"pthw.com/asymmetric-for-biometric/models"
)

type App struct {
	httpServer  *http.Server
	biometricUC biometric.UseCase
}

func NewApp() *App {
	db := initDB()
	if os.Getenv("ENV") == "debug" {
		db.Debug()
	}
	db.AutoMigrate(&models.UserBiometric{})
	db.AutoMigrate(&models.CHALLENGE{})

	biometricRepo := biometric.NewBiometricRepository(db)
	return &App{
		biometricUC: biometric.NewBiometricUseCase(biometricRepo),
	}
}

// Run gin routing https
func (a *App) Run(port string) error {

	// Init gin handler
	if os.Getenv("ENV") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// API endpoints
	biometricHttp.RegisterHTTPEndpoints(router, a.biometricUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

// Connection to db with GORM
func initDB() *gorm.DB {
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, DBHOST, DBPORT, DBNAME)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Mysql connected!!")
	}
	return db
}
