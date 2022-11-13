package main

import (
    "io/fs"
    "github.com/afanasjev/photomover"
    "time"
    "fmt"
    "path/filepath"
    "strings"
    "os"
)

func main() {

    srcDir := "/root/dev/pics"
    dst := "/root/dev/mypictures/sorted"
    if _, err := os.Stat(dst); err != nil {
        os.Mkdir(dst, 0755)
    }
    filepath.Walk(srcDir, func(srcPath string, info fs.FileInfo, err error) error {
          if !strings.HasSuffix(srcPath, ".JPG") {
              return nil
          }
          if err != nil {
              fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", srcPath, err)
              return err
          }

          if info.IsDir() {
              return nil
          } else {
              date := photomover.GetPhotoDate(srcPath)
              layout := "2006:01:02 15:04:05"
              t, err := time.Parse(layout, date)
              if err != nil {
                  fmt.Printf("on parse error: %v, %v\n", srcPath, err)
              }
              dstYear := filepath.Join(dst, fmt.Sprintf("%4d", t.Year()))
              if _, err := os.Stat(dstYear); err != nil {
                  fmt.Printf("Create dir: %v\n", dstYear)
                  os.Mkdir(dstYear, 0755)
              }
              dstMonth := filepath.Join(dstYear, fmt.Sprintf("%02d", t.Month()))
              if _, err := os.Stat(dstMonth); err != nil {
                  fmt.Printf("Create dir: %v\n", dstMonth)
                  os.Mkdir(dstMonth, 0755)
              }

              _, fileName := filepath.Split(srcPath)
              dstPath := filepath.Join(dstMonth, fileName)

              fmt.Printf("%v ---> %v\n", srcPath, dstPath)
              imageData, err := os.ReadFile(srcPath)
              if err != nil {
                  fmt.Println(err)
              }
              err = os.WriteFile(dstPath, imageData, 0644)
              if err != nil {
                  fmt.Println(err)
              }
              return nil
          }
    })
}


