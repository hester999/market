package main

import (
	"log"
	"market/app/internal/db"
	"market/app/internal/handler/ads"
	"market/app/internal/handler/auth"
	"market/app/internal/handler/image"
	"market/app/internal/handler/reg"
	authmiddle "market/app/internal/middleware/auth"
	"market/app/internal/repo/ads_repo"
	"market/app/internal/repo/auth_repo"
	"market/app/internal/repo/img_repo"
	"market/app/internal/repo/reg_repo"
	"market/app/internal/router"
	adus "market/app/internal/usecases/ads"
	authus "market/app/internal/usecases/auth"
	imgus "market/app/internal/usecases/img"
	regus "market/app/internal/usecases/reg"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "market/app/docs"
	_ "market/app/internal/handler/ads"
	_ "market/app/internal/handler/ads/dto"
	_ "market/app/internal/handler/auth"
	_ "market/app/internal/handler/auth/dto"
	_ "market/app/internal/handler/image"
	_ "market/app/internal/handler/reg"
	_ "market/app/internal/handler/reg/dto"
)

// @title Market API
// @version 1.0
// @description This is the API documentation for the Market backend.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database, err := db.Connection()
	if err != nil {
		panic(err)
	}

	adsRepo := ads_repo.NewAdsRepository(database)
	authRepo := auth_repo.NewAuthRepo(database)
	imgRepo := img_repo.NewImgRepo(database)
	regRepo := reg_repo.NewRegistry(database)

	imgUsecase := imgus.NewImgUsecase(imgRepo)
	authUsecase := authus.NewAuth(authRepo)
	adsUsecase := adus.NewAds(adsRepo, imgRepo)
	regUsecase := regus.NewRegistry(regRepo)

	imgHandler := image.NewImageHandler(imgUsecase)
	authHandler := auth.NewAuthHandler(authUsecase)
	adsHandler := ads.NewAdsHandler(adsUsecase)
	regHandler := reg.NewRegistryHandler(regUsecase)

	authMiddleware := authmiddle.AuthMiddleware(authUsecase)
	authOptionalMiddleware := authmiddle.OptionalAuth(authUsecase)

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) // Swagger UI

	app := router.NewRouter(
		authHandler,
		regHandler,
		adsHandler,
		imgHandler,
		authMiddleware,
		authOptionalMiddleware,
	)

	r.PathPrefix("/").Handler(app)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
