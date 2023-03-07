package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	filePath := flag.String("path", "", "Path to file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	dat, err := os.ReadFile(*filePath)
	check(err)
	fmt.Print(dat)
	stringified := string(dat)
	h := sha256.New()
	h.Write([]byte(stringified))
	fmt.Printf("SHA256: %x, %d", h.Sum(nil), len(dat))
	os.Setenv("FOO", fmt.Sprintf("%x", h.Sum(nil)))
	fmt.Print(os.Getenv("FOO"))

}
