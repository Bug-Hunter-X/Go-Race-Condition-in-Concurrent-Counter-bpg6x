```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        count := 0

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        count += i
                        m.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Printf("Count: %d\n", count)
}
```