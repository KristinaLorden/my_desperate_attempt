package main

import (
  "fmt"
  "log"
)

func main() {
  var a string

  for {
    fmt.Println("Vadim, do you still love me?")
    _, err := fmt.Scan(&a)
    if err != nil {
      log.Fatal("Все сломалось")
    }
    if a == "Yes" {
      fmt.Println("I love you too... I'm very sorry about what happened, and I want to talk to you tet-a-tet. Please, text me about your thoughts on telegram")
      break
    } else {
      fmt.Println("I understand you. Please let me thank you for everything that happened between us. And then I'll never bother you.")
    }
  }

}
