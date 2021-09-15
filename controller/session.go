package controller

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var CookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := CookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

// func loginHandler(response http.ResponseWriter, request *http.Request) {
// 	name := request.FormValue("name")
// 	pass := request.FormValue("password")
// 	redirectTarget := "/"
// 	if name != "" && pass != "" {
// 		// .. check credentials ..
// 		SetSession(name, response)
// 		redirectTarget = "/internal"
// 	}
// 	http.Redirect(response, request, redirectTarget, 302)
// }

// logout handler

// func logoutHandler(response http.ResponseWriter, request *http.Request) {
// 	ClearSession(response)
// 	http.Redirect(response, request, "/", 302)
// }
