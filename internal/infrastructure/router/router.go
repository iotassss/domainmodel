package router

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/domainmodel/internal/handler"
	"github.com/iotassss/domainmodel/internal/infrastructure/database"
	"github.com/iotassss/domainmodel/internal/presenter"
	"github.com/iotassss/domainmodel/internal/repository"
	"github.com/iotassss/domainmodel/internal/repository/model"
	"github.com/iotassss/domainmodel/internal/usecase"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	db, err := database.NewDB()
	if err != nil {
		slog.Error("database connection failed", slog.Any("error", err))
		os.Exit(1)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Credential{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")

	userRepository := repository.NewUserRepository(db)
	createUserPresenter := presenter.NewAPICreateUserPresenter()
	createUserInteractor := usecase.NewCreateUserInteractor(userRepository, createUserPresenter)
	getUserByUUIDPresenter := presenter.NewAPIGetUserByUUIDPresenter()
	getUserByUUID := usecase.NewGetUserByUUIDInteractor(userRepository, getUserByUUIDPresenter)
	userHandler := handler.NewUserHandler(createUserInteractor, getUserByUUID)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	r.GET("/db/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to get DB instance"})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "DB not ready"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "DB is ready"})
	})

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/users", userHandler.CreateUser)
		apiRouter.GET("/users/:uuid", userHandler.GetUserByUUID)
	}

	return r
}
