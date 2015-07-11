package f2k

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/pod", pod_handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "f2k(Feed2Knowledge) is an API to provide one HN Submission a day. Currently the information is sourced from HN")
}

func pod_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hola")
}
