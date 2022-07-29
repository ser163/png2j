# png2j
Convert PNG to JPG GO library

## Description

    png2j is a Go library to convert PNG to JPG.
## Install
```shell
  go get -u github.com/jteeuwen/png2j
```     

## Usage

```go
  
    package main

    import "github.com/ser163/png2j"

    func main() {
       // Convert PNG to JPG
        err := png2j.Con2jpg("./image/go.png", "./image/go.png.jpg")
        if err != nil {
            panic(err)
        }
        // Resize JPG image
        err2 := png2j.ReSizeImage("./image/go.png.jpg", 948, 418, "./image/go_big.png.jpg")
        if err2 != nil {
            panic(err2)
        }
    }
```

       
   
   