package config

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
	biometric "pthw.com/asymmetric-for-biometric/internal/biometric"
)

var (
	db *gorm.DB
)

type App struct {
	httpServer *http.Server

	biometricUC biometric.UseCase
}

func NewApp() *App {
	db := initDB()

	// userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	// bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	// return &App{
	// 	bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
	// 	authUC: authusecase.NewAuthUseCase(
	// 		userRepo,
	// 		viper.GetString("auth.hash_salt"),
	// 		[]byte(viper.GetString("auth.signing_key")),
	// 		viper.GetDuration("auth.token_ttl"),
	// 	),
	// }

	biometricRepo := biometric.NewBiometricRepository(db)
	return &App{
		biometricUC: biometric.NewBiometricUseCase(biometricRepo),
	}

}

func (a *App) Run(port string) error {
	// Init gin handler
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

func initDB() *gorm.DB {
	dsn := "dbeaver:dbeaver@tcp(localhost:3306)/showcase_db?charset=utf8mb4&parseTime=True&loc=Local"
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
