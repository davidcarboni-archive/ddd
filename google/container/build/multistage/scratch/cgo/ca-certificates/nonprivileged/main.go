// https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get("https://google.com")
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Printf("Length of content received from google.com: %d\n", len(body))
    fmt.Printf("Running as user ID: %d\n", os.Getuid())
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
