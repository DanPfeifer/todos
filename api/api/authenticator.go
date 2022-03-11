package api

import (
	"github.com/pkg/errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"todos/resources"
)
var jwtKey = []byte("123456")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


// Define our struct
type authenticator struct {
	d UserStore
}

func NewAuthenticator(d UserStore)(a *authenticator) {
	return &authenticator{
		d: d,
	}
}

func validateToken( req *http.Request) (claims *Claims, err error) {
	token := req.Header.Get("X-Session-Token")
	if token == "" {
		err = errors.New("Unauthorized")
		return
	}
	claims = &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = errors.New("Unauthorized")
			return
		}
		err = errors.New("Unauthorized")
		return
	}
	if !tkn.Valid {
		err = errors.New("Unauthorized")
		return
	}
	return
}

func (a *authenticator) GetUser(req *http.Request) (user resources.User, err error) {
	claims, err := validateToken(req)
	if err != nil {
		err = errors.New("Unauthorized")
		return
	}
	user, err = a.d.FindByEmail(claims.Username)
	return
}

func (a *authenticator) GenerateToken(username string) (tokenString string, err error){
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	if err != nil {
		return
	}
	return
}