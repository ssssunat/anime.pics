package main

import (
	"encoding/json"
	"fmt"
    "os"
    "io"
	"io/ioutil"
	"log"
	"net/http"
    "os/exec"
    "runtime"
)

type urlResponse struct {
    URL string      `json:"url"`  
}
// lets gooooooooo
func main() {

    resp, err := http.Get("https://api.waifu.pics/sfw/kiss")
    
    if err != nil {
        log.Fatal(err)
    }
    
    defer resp.Body.Close()

    fmt.Println("Response status: ", resp.StatusCode)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    var r urlResponse
    if err = json.Unmarshal(body, &r); err != nil {
        log.Fatal(err)
    }

    url := r.URL
 
    // Make an HTTP GET request to the URL
    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to download image:", err)
        return
    }
    defer response.Body.Close()

    // Save the image to a local file
    file, err := os.Create("image.jpg")
    if err != nil {
        fmt.Println("Failed to create image file:", err)
        return  
    }

    defer file.Close()

    _, err = io.Copy(file, response.Body)
    if err != nil {
        fmt.Println("Failed to save image:", err)
        return  
    }

    fmt.Println("Image downloaded and saved successfully.")

    // Open the image file with the default image viewer on the system
    err = OpenImage("image.jpg")
    if err != nil {
        fmt.Println("Failed to open image:", err)
        return
    }
}

// OpenImage opens an image file with the default image viewer on the system
func OpenImage(filename string) error {
    var cmd *exec.Cmd
    switch runtime.GOOS {
        case "darwin":
            cmd = exec.Command("open", filename)
        case "linux":
            cmd = exec.Command("xdg-open", filename)
        case "windows":
            cmd = exec.Command("cmd", "/c", "start", filename)
        default:
            return fmt.Errorf("unsupported platform")
    }
    return cmd.Start()
}