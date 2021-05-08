package main

import ( 
  "io/ioutil"
  "log"
  "sync"
  "github.com/fatih/color"
  "flag"
)

var wait sync.WaitGroup

func Show() {
  files, err := ioutil.ReadDir("./")
  


  if err != nil {
    log.Fatal(err)
  }

  for i, f := range files {
    color.Cyan("%d %s",i, f.Name())
  }

  wait.Done()
}

func main() {
  wait.Add(1)
  go Show()
  wait.Wait()
}
