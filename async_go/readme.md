## What is GoRoutine?
Independently executed function.

```go
package main

import "fmt"

func main() {
	go hello()
}
func hello() {
	fmt.Println("It's most likely you will never see this.")
}
```
This code show nothing, why?
Go Concurrency:

- Go is a concurrent language, meaning it can execute multiple tasks simultaneously (goroutines) even within a single program.
    In your code, the main() function starts a goroutine by calling go hello(). This means the hello() function runs asynchronously, parallel to the main() function.
    The main() function then finishes its execution without waiting for hello() to complete. If the main() function doesn't have any other statements after starting the goroutine, it will exit immediately.

Output:

- Since the main() function exits before hello() finishes, there wouldn't be any code left to print the message. Therefore, if you run the program as-is, you might not see the output.

## Goroutines and Channels in Go: A Detailed Explanation

In Go, two key concepts power concurrent programming: **goroutines** and **channels**. Let's delve into each:

**Goroutines:**

* **What they are:** Goroutines are lightweight threads of execution. They run concurrently within a single process, allowing your program to perform multiple tasks **simultaneously**. Think of them as independent pieces of code that can run in the background while your main program continues.
* **Benefits:** Goroutines offer several advantages:
    * **Scalability:** They handle concurrency efficiently, allowing you to scale your application's performance without creating heavy operating system threads.
    * **Responsiveness:** The main program doesn't block when a goroutine starts, so your application remains responsive to user input.
    * **Resource efficiency:** Goroutines use minimal memory compared to traditional threads, making them more lightweight.
* **Creating goroutines:** You create a goroutine using the `go` keyword followed by a function call like this:

```go
go myFunc()
```

This tells Go to run the `myFunc()` function as a goroutine in the background.

**Channels:**

* **What they are:** Channels are a communication mechanism between goroutines. They act like pipes or tubes that allow goroutines to **send and receive data** to each other.
* **Benefits:** Channels enable safe and structured communication:
    * **Synchronization:** You can use channels to synchronize execution between goroutines, ensuring they complete tasks in a specific order or exchange data safely.
    * **Data buffering:** Channels can act as buffers, allowing goroutines to send data at different speeds without losing information.
* **Types of channels:** There are two main types of channels:
    * **Unbuffered:** These channels hold no data and block the sender until a receiver is ready.
    * **Buffered:** These channels can hold a certain amount of data, allowing the sender to proceed without waiting for a receiver (up to the buffer size).
* **Using channels:** You create channels using the `make` function with the channel type as an argument. Sending and receiving data is done with the `<-` operator:

```go
ch := make(chan int) // Create an unbuffered channel
ch <- 42 // Send value 42 to the channel
value := <-ch // Receive value from the channel
```

**Understanding the Relationship:**

Goroutines and channels work together to power concurrent programming in Go. Goroutines provide the execution units, while channels facilitate communication and coordination between them.

**Example:**

Here's a simple example demonstrating goroutines and channels:

```go
package main

import (
    "fmt"
)

func worker(ch chan int) {
    for i := 0; i < 5; i++ {
        fmt.Println("Worker:", i)
        ch <- i // Send value to channel
    }
}

func main() {
    ch := make(chan int)
    go worker(ch) // Start worker goroutine

    for i := 0; i < 3; i++ {
        value := <-ch // Receive value from channel
        fmt.Println("Main:", value)
    }
}
```

In this example, the `worker` goroutine sends values to the channel, while the `main` goroutine receives them. They run concurrently but communicate and synchronize through the channel.

**Conclusion:**

Goroutines and channels are fundamental building blocks for concurrency in Go. By understanding their concepts and applications, you can write efficient, scalable, and responsive programs that take advantage of Go's powerful concurrency features.

## Go doesn't directly use OS threads.
You're correct! Go **doesn't directly use operating system threads** for its goroutines. Instead, it relies on a lightweight thread management system within the Go runtime itself. Here's a breakdown of the key differences:

**Operating System Threads:**

* Heavy and resource-intensive.
* Each thread requires a significant amount of memory and system resources.
* Creating and managing threads can be expensive.
* Not easily scalable due to resource limitations.

**Go Goroutines:**

* Lightweight and efficient.
* Use a small amount of memory and system resources compared to threads.
* Much faster to create and manage.
* Can be scaled more easily due to their lightweight nature.

However, it's important to understand how Go achieves this efficiency:

* **Multiplexing:** Go relies on a technique called **multiplexing**, where multiple goroutines share a small number of actual OS threads. This allows the Go runtime to efficiently manage a large number of goroutines without needing a dedicated thread for each one.
* **Scheduling:** Go runtime has a sophisticated scheduling system that determines which goroutines to run on available OS threads. This ensures efficient utilization of resources and good overall performance.
* **Cooperative Multitasking:** Goroutines themselves don't preemptively switch execution like threads. Instead, they rely on cooperative multitasking, where a goroutine explicitly yields control back to the runtime when it needs to block or perform other tasks.

In summary, while Go doesn't directly use OS threads for each goroutine, it provides an efficient system for managing them using a combination of multiplexing, scheduling, and cooperative multitasking. This allows Go to offer the benefits of concurrency without the resource overhead of traditional threads.

Additionally, it's worth noting that Go does offer other concurrency primitives like channels and synchronization mechanisms that enable you to control and manage how goroutines interact, providing a powerful and flexible approach to building concurrent applications.