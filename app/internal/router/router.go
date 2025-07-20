package router

import (
	"github.com/gorilla/mux"
	"market/app/internal/handler/ads"
	"market/app/internal/handler/auth"
	"market/app/internal/handler/image"
	"market/app/internal/handler/reg"
	"net/http"
)

func NewRouter(
	authHandler *auth.AuthHandler,
	regHandler *reg.RegistryHandler,
	adsHandler *ads.AdsHandler,
	imageHandler *image.ImageHandler,
	authMiddleware func(http.Handler) http.Handler,
	authOptionalMiddleware func(http.Handler) http.Handler,
	imgMiddleware func(http.Handler) http.Handler,
) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	api.HandleFunc("/register", regHandler.RegistrationHandler).Methods(http.MethodPost)
	api.HandleFunc("/login", authHandler.LoginHandler).Methods(http.MethodPost)
	api.Handle("/logout", authMiddleware(http.HandlerFunc(authHandler.Logout))).Methods(http.MethodPost)

	// Ads
	api.Handle("/ads", authMiddleware(http.HandlerFunc(adsHandler.Create))).Methods(http.MethodPost)
	api.Handle("/ads", authOptionalMiddleware(http.HandlerFunc(adsHandler.GetAllAds))).Methods(http.MethodGet)
	api.HandleFunc("/ads/{id}", adsHandler.GetAdByID).Methods(http.MethodGet)
	api.Handle("/ads/{id}", authMiddleware(http.HandlerFunc(adsHandler.Delete))).Methods(http.MethodDelete)

	// Images
	api.Handle("/ads/{id}/images", authMiddleware(http.HandlerFunc(imageHandler.AddImage))).Methods(http.MethodPost)
	api.HandleFunc("/ads/{id}/images", imageHandler.GetImages).Methods(http.MethodGet)
	api.HandleFunc("/ads/images/{id}", imageHandler.GetImageById).Methods(http.MethodGet)

	return r
}
