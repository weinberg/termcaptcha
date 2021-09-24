package main

import (
  "github.com/brickshot/termcaptcha/internal"
  "os"
  "time"
)

func main() {
  // _, c := internal.OneNumber()
  _, c := internal.OneLine()
  for _,x := range c {
    os.Stdout.Write([]byte{byte(x)})
    time.Sleep(1000000)
  }
  os.Stdout.Write([]byte("\n"))
}
