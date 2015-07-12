package main

import (
  "fmt"
  "net/http"
  "math/rand"
  "time"
  "strconv"
  "encoding/json"
)

type Results struct {
  Result []Story `json:"hits"`
}

type Story struct {
  Title, Url string
}

func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  http.HandleFunc("/", handler)
  http.HandleFunc("/pod", pod_handler)

  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "f2k(Feed2Knowledge) is an API to provide one HN Submission a day. Currently the information is sourced from HN")
}

func pod_handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var results Results

  hn_url := "http://hn.algolia.com/api/v1/"
  page := rand.Intn(50)
  points := randomInt(200, 600)
  url := hn_url + "search?tags=story&numericFilters=points>" + strconv.Itoa(points) + "&page=" + strconv.Itoa(page)

  resp, err := makeRequest(url)

  if err != nil {
    fmt.Fprint(w, "Something Went Wrong!")
  }

  err = json.NewDecoder(resp.Body).Decode(&results)

  if err != nil {
    fmt.Fprint(w, "Something Went Wrong!")
  } else {
    res, _ := json.Marshal(results.Result[0])
    fmt.Fprint(w, string(res))
  }

}

func randomInt(min, max int) int {
  return min + rand.Intn(max - min)
}

func makeRequest(url string) (*http.Response, error) {
  response, err := http.Get(url)
  if err != nil {
    return nil, err
  }


  if response.StatusCode == http.StatusNotFound {
    return nil, fmt.Errorf(http.StatusText(http.StatusNotFound))
  }
  return response, nil
}


