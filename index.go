package main

import "net/http"

func main() {
	http.HandleFunc("/signup", signupPage)
	// http.HandleFunc("/login", loginPage)
	// http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}
}
