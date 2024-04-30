package auth

import (
	"fmt"
	"net/http"
	"time"
)

// Struct Section
type Session struct {
	username string
	expiry   time.Time
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func (s Session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func AuthenticationHandlerester() {
	fmt.Println("auth is here")
}

func Auth_check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the auth check!")
}

func Signin(w http.ResponseWriter, r *http.Request) {

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome!")
}
