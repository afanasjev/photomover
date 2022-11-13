package main

import (
    "github.com/afanasjev/photomover"
    "fmt"
    "time"
)

func main() {
    filepath := "./IMG_0148.JPG"
    date := photomover.GetPhotoDate(filepath)
    layout := "2006:01:02 15:04:05"
    t, err := time.Parse(layout, date)

    if err != nil {
          fmt.Println(err)
    }
    fmt.Println(t.Day())
    fmt.Println(t.Month())
    fmt.Println(t.Year())
}


