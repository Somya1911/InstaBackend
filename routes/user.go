package routes

import (
	
	"net/http"
	"fmt"
	"/Users/somyakhatri/InstaBackend/models /user.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
 func


)




func postRequests() {
 http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})}

