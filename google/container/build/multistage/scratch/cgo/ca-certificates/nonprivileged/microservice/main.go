// https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
// http://devcenter.wercker.com/docs/quickstarts/advanced/building-minimal-containers-with-go

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "log"
)

type ContentLengthResponse struct {
    Google int `json:"google"`
}

func ContentLengthHandler(res http.ResponseWriter, req *http.Request) {
    resp, err := http.Get("https://google.com")
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Printf("Length of content received from google.com: %d\n", len(body))
    fmt.Printf("Running as user ID: %d\n", os.Getuid())
    contentLengthResponse := &ContentLengthResponse{
        Google : len(body),
    }
    data, _ := json.MarshalIndent(contentLengthResponse, "", "  ")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    log.Println("Listening on this host: http://localhost:7000")

    http.HandleFunc("/google", ContentLengthHandler)
    err := http.ListenAndServe(":7000", nil)
    if err != nil {
        log.Fatal("Unable to listen on :7000: ", err)
    }
}
