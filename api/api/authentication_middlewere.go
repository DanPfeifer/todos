package api

import (
	"net/http"
	"todos/resources"
)

type Authenticator interface {
	GetUser(r *http.Request) (user resources.User, err error)
	GenerateToken(username string) (tokenString string, err error)
}

type authMiddleware struct {
	auth Authenticator
}

func NewAuthenticationMiddleware(a Authenticator) *authMiddleware {
	return &authMiddleware{
		auth: a,
	}
}

// Middleware function, which will be called for each request
func (amw *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, err := amw.auth.GetUser(req)
		if err != nil {
			jsonErr(res, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(res, req)
	})
}

