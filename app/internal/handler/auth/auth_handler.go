package auth

type Auth interface {
	Login(email, password string) (string, error)
	Logout(token string) error
	ValidateSession(token string) (string, error)
}

type AuthHandler struct {
	usecase Auth
}

func NewAuthHandler(autUsecases Auth) *AuthHandler {
	return &AuthHandler{
		usecase: autUsecases,
	}
}
