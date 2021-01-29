package main

import (
    "fmt"
    "runtime"
    "syscall/js"
    "time"
)

func rewrite() {
    t := time.Now()
    document := js.Global().Get("document")
    element := document.Call("getElementById", "text")
    element.Set("innerText", t.String())
}

func main() {
    fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
    for range time.Tick(1 * time.Millisecond) {
        go rewrite()
    }
}
