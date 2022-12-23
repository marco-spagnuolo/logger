 package logger

 import (
  "fmt"
 )
    
 var Version string = "1.0"
    
 func Log(mess string) {
  fmt.Println("[LOG] " + mess)
 }
