A **CORS error** happens because your browser blocks requests from one origin (your frontend) to another (your Go server) unless the server explicitly allows it.

You need to add **CORS headers in your Go server response**.

---

# 1️⃣ Simple CORS fix in Go (manual headers)

Add these headers in your handler:

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	fmt.Fprintln(w, "Hello from Go server")
}

func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8080", nil)
}
```

---

# 2️⃣ If using fetch

Example frontend request:

```javascript
fetch("http://localhost:8080/api")
  .then(res => res.text())
  .then(data => console.log(data));
```

---

# 3️⃣ Proper CORS middleware (recommended)

Instead of repeating headers everywhere:

```go
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handler)

	http.ListenAndServe(":8080", enableCors(mux))
}
```

---

# 4️⃣ Production-safe version

Instead of `"*"` allow only your frontend:

```go
w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
```

---

# 5️⃣ Common mistake with `fetch`

If you send JSON:

```javascript
fetch("http://localhost:8080/api", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({name: "John"})
});
```

Then your Go server **must allow**:

```go
Access-Control-Allow-Headers: Content-Type
```

---

✅ **Quick summary**

Add these headers in Go:

```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

---

💡 If you want, I can also show:

* **The cleanest CORS middleware used in production Go servers**
* **How to fix CORS in Go + React / Go + vanilla JS**
* **Why browsers send a preflight OPTIONS request** (this confuses many developers).

