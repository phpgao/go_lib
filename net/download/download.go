package main

import (
	"io"
	"net/http"
	"os"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	downloadFile("/tmp/PhpStorm-2016.3.2.dmg", "https://download.jetbrains.8686c.com/webide/PhpStorm-2016.3.2.dmg")
}
