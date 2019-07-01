package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

func runWego() {
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
		runWego()
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		runWego()
	}
	matched, _ := regexp.Match("^[<>A-Za-z0-9]", contents)
	if matched {
		runWego()
	} else {
		fmt.Printf("%s", contents)
	}
}
