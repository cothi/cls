package main

import (
	"color"
  "flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

var wait sync.WaitGroup

func Show(Location string) {
  files, err := ioutil.ReadDir(Location)
  
  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    fmt.Printf(color.White, f)
  }

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
