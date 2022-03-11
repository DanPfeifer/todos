package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"todos/resources"
)

type UserStore interface {
	Login(email string, password string) (u resources.User, err error)
    FindByEmail(email string) (u resources.User, err error)
    SetToken(user resources.User, token string) (u resources.User, err error)
}

type UserResponse struct {
    ID int `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Token string `json:"token"`
}

type LoginRequest struct {
   Email string `json:"email"`
   Password string `json:"password"`
}

type userHandler struct {
    d UserStore
    auth Authenticator
}

func NewUserHandler(db UserStore, auth Authenticator) *userHandler {
	return &userHandler{
		d: db,
		auth: auth,
	}
}

func (uh *userHandler) login(res http.ResponseWriter, req *http.Request) {

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        jsonErr(res, "Forbidden", http.StatusForbidden)
        return
    }

    var login LoginRequest
    if err := json.Unmarshal(body, &login); err != nil {
        jsonErr(res, "Forbidden", http.StatusForbidden)
        return
    }

    user, err := uh.d.Login(login.Email, login.Password)
    if err != nil {
        jsonErr(res, "Forbidden", http.StatusForbidden)
        return
    }

    token, err := uh.auth.GenerateToken(user.Email)
    if err != nil {
        jsonErr(res, "Forbidden", http.StatusForbidden)
        return
     }

    user, err = uh.d.SetToken(user, token)

    if err != nil {
    fmt.Print(err)
        jsonErr(res, "Forbidden", http.StatusForbidden)
        return
    }

    resBody :=  mapUserResponse(user)
    writeJson(res, resBody, http.StatusOK)
}

func (uh *userHandler) showUser(res http.ResponseWriter, req *http.Request) {
    user, err := uh.auth.GetUser(req)

	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusBadRequest)
		return
	}

	resBody := mapUserResponse(user)

	writeJson(res, resBody, http.StatusOK)
}

func mapUserResponse(u resources.User) UserResponse {
    return UserResponse{
        ID: u.ID,
        FirstName: u.FirstName,
        LastName: u.LastName,
        Email: u.Email,
        Token: u.Token,
    }
}
