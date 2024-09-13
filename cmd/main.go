package main

import (
	"backend-carpintery/internal/repositories/postgres"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	collageHandler "backend-carpintery/cmd/api/handlers/collage"
	collageRepo "backend-carpintery/internal/repositories/postgres/collage"
	collageService "backend-carpintery/internal/services/collage"

	productHandler "backend-carpintery/cmd/api/handlers/product"
	productRepo "backend-carpintery/internal/repositories/postgres/product"
	productService "backend-carpintery/internal/services/product"

	bannerHandler "backend-carpintery/cmd/api/handlers/banner"
	bannerRepo "backend-carpintery/internal/repositories/postgres/banner"
	bannerService "backend-carpintery/internal/services/banner"

	categoryHandler "backend-carpintery/cmd/api/handlers/categories"
	categoryRepo "backend-carpintery/internal/repositories/postgres/categories"
	categoryService "backend-carpintery/internal/services/categories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el env file: %w", err)
	}
	ginEngine := gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	portAsNumber, _ := strconv.Atoi(dbPort)

	connectkey := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", portAsNumber, dbUser, dbPassword, dbName)
	db, err := postgres.CreateConnection(connectkey)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Products Dependencies
	collageRepo := collageRepo.Repository{
		Db: db}
	collageSrv := collageService.Services{
		Repo: collageRepo}
	collageHandler := collageHandler.Handler{
		CollageService: collageSrv}

	//Products Dependencies
	productRepo := productRepo.Repository{
		Db: db}
	productSrv := productService.Services{
		Repo: productRepo}
	productHandler := productHandler.Handler{
		ProductService: productSrv}

	//Banner Dependencies
	bannerRepo := bannerRepo.Repository{
		Db: db}
	bannerSrv := bannerService.Services{
		Repo: bannerRepo}
	bannerHandler := bannerHandler.Handler{
		BannerService: bannerSrv}

	//Categories Dependencies
	categoryRepo := categoryRepo.Repository{
		Db: db}
	categorySrv := categoryService.Services{
		Repo: categoryRepo}
	categoryHandler := categoryHandler.Handler{
		CategoryService: categorySrv}

	//Collage Routes
	ginEngine.POST("/create-collage", collageHandler.CreateCollage)
	//Banner Routes
	ginEngine.POST("/create-banner", bannerHandler.CreateBanner)
	//Card Routes
	ginEngine.POST("/create-product", productHandler.CreateProduct)
	//Category Routes
	ginEngine.POST("/create-category", categoryHandler.CreateCategory)

	log.Println("Servidor corriendo en el puerto 8080")
	ginEngine.Run(":8080")

}
