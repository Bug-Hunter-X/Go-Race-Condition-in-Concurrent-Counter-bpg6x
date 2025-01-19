# Go Race Condition Example

This repository demonstrates a common race condition bug in Go programs involving concurrent access to a shared variable without proper synchronization.  The `bug.go` file contains the buggy code, and `bugSolution.go` provides a corrected version.

The bug arises when multiple goroutines attempt to update the `count` variable simultaneously.  Without a mutex or other synchronization mechanism, the final value of `count` will be incorrect and non-deterministic.

The solution uses a `sync.Mutex` to protect the `count` variable, ensuring that only one goroutine can access and modify it at a time.