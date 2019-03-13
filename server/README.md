### Example

```go
// Exemple middleware
func GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Example get body
		username := r.FormValue("username")
    log.Println(username)
    // Example get Header
    ua := r.Header.Get("User-Agent")
    log.Println(ua)
    // Example get params (/{id}/)
    vars := mux.Vars(request)
    id := vars["id"]

		next.ServeHTTP(w, r)
	})
}
```
