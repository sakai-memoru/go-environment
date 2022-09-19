package main

import (
	"fmt"
)

func doMain(){
  // ------------------- begin.
  println("// --------------------- begin.")
  
  // ------------------- process.
  name := "Golang"
  // var name string = "Golang"
  // var name = "Golang"
  fmt.Printf("Hello, %v !!\n", name)
  
  // ------------------- end.
  println("// --------------------- end...")
}

func main() {
  println("-- ---------------------")
	doMain()
  
}
