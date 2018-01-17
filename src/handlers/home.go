package handlers

import(
  "fmt"
  "net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
  fmt.Fprint(w, "Hello World");
}
