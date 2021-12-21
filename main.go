package main 

import (
    "fmt"
    "sync"
)

func main() {
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 10; i++ {
	fmt.Println("done")
	fmt.Println("feature/f2")
	fmt.Println("create develop")
	fmt.Println("log3")
        <-done
    }
}
