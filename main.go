package main

import (
  "net/http"
  "io/ioutil"
  "os"
  "os/exec"
  "fmt"
  "regexp"
)

func run_wego() {
  out, err := exec.Command("wego").Output()
  if err != nil {
    fmt.Printf("%v", err)
    os.Exit(1)
  }

  fmt.Printf("%s", out)
}

func main() {
  response, err := http.Get("http://wttr.in?format=2")
  if err != nil {
    run_wego()
  }
  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    run_wego()
  }
  matched, _ := regexp.Match("^[A-Za-z0-9]", contents)
  if matched {
    run_wego()
  } else {
    fmt.Printf("%s", contents)
  }
}
