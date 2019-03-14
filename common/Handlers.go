package common

import (
	"awesomeProject/login-reg-cookie-demo/helpers"
	"awesomeProject/login-reg-cookie-demo/repos"
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Handlers

// for GET
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, _ = helpers.LoadFile("templates/login.html")
	fmt.Fprintf(w, body)
}

// for POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")

	redirectTarget := "/"

	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		// Database check for user data
		_userIsValid := repos.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, w)
			redirectTarget += "index"
		} else {
			redirectTarget += "register"
		}
	}

	http.Redirect(w, r, redirectTarget, 302)
}

// for GET
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, _ = helpers.LoadFile("templates/register.html")
	fmt.Fprintf(w, body)
}

// for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !helpers.IsEmpty(uName)
	_email = !helpers.IsEmpty(email)
	_pwd = !helpers.IsEmpty(pwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _uName && _email && _pwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

// for GET
func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !helpers.IsEmpty(userName) {
		var indexBody, _ = helpers.LoadFile("templates/index.html")
		fmt.Fprintf(w, indexBody, userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// for POST
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearCookie(w)
	http.Redirect(w, r, "/", 302)
}

// Cookie

func SetCookie(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ClearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func GetUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
