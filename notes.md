# Notes from a beginner golang developer

## Testing sql queries in golang

Testing sql queries is a thing. When I test my API handlers, at some point I need to be able to test against a data source. In that case, I see two main options:

- mocking databases -> most popular option: [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock?source=post_page-----5af19075e68e--------------------------------)
- testing against a real database

Mocking database is most of the time an overkill and pretty cumbersome, even for big organisations. Testing against a real database is a good option since spinning up a database nowadays is just spinning up a container. As containers are easy to destruct, it is a solution for db testing.

There are some tools out there to spin up docker containers easily in go code. [dockertest](https://github.com/ory/dockertest) is one of those.

Read [this post](https://www.reddit.com/r/golang/comments/u62emg/mocking_database_or_use_a_test_database/) for the whole argument around this topic.

## TestMain and the use case of it

> More details in [this article](https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc).

TestMain will be run instead of running the tests directly. The M struct contains methods to access and run the tests.

There is a single TestMain for each package.

`testing.M` has a single defined function named `Run`. It runs all the tests and returns an exit code to be passed to `os.Exit`

```go
package mypackagename

import (
    "testing"
    "os"
)

func TestMain(m *testing.M) {
    log.Println("Do stuff BEFORE the tests!")
    exitVal := m.Run()
    log.Println("Do stuff AFTER the tests!")

    os.Exit(exitVal)
}

func TestA(t *testing.T) {
    log.Println("TestA running")
}

func TestB(t *testing.T) {
    log.Println("TestB running")
}
```

## Panic, log.Fatal and recover

### panic

A typical use case for panic is when your program encounters a situation where continuing execution would lead to inconsistent or undefined behavior. For example, trying to access an index beyond the bounds of an array, or attempting to dereference a nil pointer can cause a panic.

```go
package main

import "fmt"

func main() {
    // Simulating a panic by dividing by zero
    result := 10 / 0
    fmt.Println("Result:", result) // This line won't be reached
}
```

### log.Fatal

Unlike panic, `log.Fatal` is typically used for situations where the program encounters an error that it cannot recover from, but it's not necessarily an unexpected or catastrophic failure. For example, failing to open a required file, or encountering invalid user input might warrant using log.Fatal.

```go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.Open("non_existent_file.txt")
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }
    // Do something with the file
    _ = file.Close()
}
```

In summary, panic is used for unexpected, unrecoverable errors that indicate a bug in the program, while log.Fatal is used for expected errors that the program cannot recover from, but are not necessarily indicative of a bug.

### recover

Recover is used with panic. If program panics, deferred function will be run. And in that deferred function you can recover the panic value:

```go
package main

import "fmt"

func recoverExample() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}

func main() {
    defer recoverExample()

    fmt.Println("Start")
    panic("Something went wrong!") // This will cause a panic
    fmt.Println("End") // This line won't be executed
}
// Output:
// Start
// Recovered: Something went wrong!
```
