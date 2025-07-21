package auth

type AuthHandler struct {
	usecase Auth
}

func NewAuthHandler(autUsecases Auth) *AuthHandler {
	return &AuthHandler{
		usecase: autUsecases,
	}
}
