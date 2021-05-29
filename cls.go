package main

import (
	"flag"
  "strings"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"github.com/wooongdev/cls/color"
  "unicode/utf8"
)



var wait sync.WaitGroup
var width1, width2 int = 12, 12



func Show(Location string) {
  files, err := ioutil.ReadDir(Location)
  
  if err != nil {
    log.Fatal(err)
  }

  var DirName []string
  var FileName []string
  var num int = 1

  for _, f := range files {
    
    if f.IsDir() {
      DirName = append(DirName, f.Name())
      if utf8.RuneCountInString(f.Name()) > width1 {
        width1 = utf8.RuneCountInString(f.Name())
      }
    } else {
      if utf8.RuneCountInString(f.Name()) > width2 {
        width2 = utf8.RuneCountInString(f.Name())
      }
      FileName = append(FileName, f.Name())
    }
  }
  println(width1, width2)
  fmt.Println("NO  FILENAME", strings.Repeat(" ", width1-9), "KINDS      NO  FILENAME")

  var long int = 0
  if len(DirName) >= len(FileName) {
    long = len(DirName)
  } else {
    long = len(FileName)
  }

  for i := 0; i < long; i++ {
    if (i < len(DirName)) {
      print(string(color.ColorRed))
      print(num, strings.Repeat(" ", 3 - len(strconv.Itoa(num))))
      print(" ",  DirName[i], strings.Repeat(" ", width1 - utf8.RuneCountInString(DirName[i]))," Directory  ",string(color.ColorReset))
    } else {
      print(strings.Repeat(" ", 16 + width1))
    }


    if (i < len(FileName)) {
      print(string(color.ColorBlue))
      print(num, strings.Repeat(" ", 3 - len(strconv.Itoa(num))))
      print(" ",  FileName[i], strings.Repeat(" ", width2 - utf8.RuneCountInString(FileName[i])), string(color.ColorReset))
      println("")
      num += 1
    } else {
      println("")
      num += 1
    }
  }

  /*
    fmt.Println(string(colorRed), "test")
    fmt.Println(string(colorGreen), "test")
    fmt.Println(string(colorYellow), "test")
    fmt.Println(string(colorBlue), "test")
    fmt.Println(string(colorPurple), "test")
    fmt.Println(string(colorWhite), "test")
    fmt.Println(string(colorCyan), "test", string(colorReset))
    fmt.Println("next")
  */ 
  wait.Done()
}


func main() {
  location := flag.String("l", "./", "a string")

  flag.Parse()

  fmt.Println(*location)
  wait.Add(1)
  go Show(*location)
  wait.Wait()

}

