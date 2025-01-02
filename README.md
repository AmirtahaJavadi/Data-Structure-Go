# 🏗️ Go Data Structures - Stack, Queue, Linked Lists & More!

Welcome to the **Go Data Structures** project! 🎉 This is where the magic happens — we bring you powerful, efficient, and simple-to-understand implementations of classic data structures in Go. Whether you're a beginner or an experienced developer, this repository is your playground to explore stacks, queues, singly linked lists, and doubly linked lists. 🚀

Ready to dive into the world of data structures? Let’s go! 🌟

## ✨ What’s Inside?

In this project, you'll find the following exciting data structures:

- **Stack**: A LIFO (Last In, First Out) powerhouse for handling data in reverse order! 🏋️‍♂️
- **Queue**: A FIFO (First In, First Out) mechanism for orderly processing of tasks. ⏳
- **Singly Linked List**: A chain of nodes connected in one direction – lightweight and efficient! 🔗
- **Doubly Linked List**: A supercharged linked list with connections in both directions, ready to rock! 🔄

## 📚 Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Data Structures Explained](#data-structures-explained)
  - [Stack](#stack)
  - [Queue](#queue)
  - [Singly Linked List](#singly-linked-list)
  - [Doubly Linked List](#doubly-linked-list)
- [Contributing](#contributing)
- [License](#license)

## 🚀 Installation <a name="installation"></a>

Let's get this project running on your machine! Clone the repo and fire it up:

```bash
git clone https://github.com/yourusername/go-datastructures.git
cd go-datastructures
cd "singly linked list"
```

Then, build and run it using Go:

```bash
go run main.go node.go linkedlist.go
```

You’re all set! 🏁 Now you can start exploring and experimenting with different data structures!

## 📦 Usage <a name="usage"></a>

### 🏋️‍♂️ Stack: Push, Pop, Repeat!

Stacks are awesome when you need a LIFO (Last In, First Out) structure. Think of it as a stack of plates: you can only add or remove from the top! 🔝
```go
stack := Stack{}

stack.Push(1)
stack.Push(2)
stack.Push(3)
stack.Print()
popped, err := stack.Pop()
stack.Print()
```
Boom! That’s a stack in action. 💥
### ⏳ Queue: FIFO, Because Order Matters!

A Queue follows the FIFO (First In, First Out) principle. It’s like a line at your favorite coffee shop: the first person to arrive gets served first! ☕️

```go
q := Queue{}
queue.Enqueue(10)  // Add 10 to the back
queue.Enqueue(20)  // Add 20 to the back
queue.Enqueue(30)  // Add 30 to the back
q.print()
q.Dequeue()
q.print()
fmt.Println(q.IsEmpty())
```
Now that's a perfectly ordered queue! 🔄

### 🔗 Singly Linked List: A One-Way Journey!
The Singly Linked List is a chain of nodes, each with data and a pointer to the next node. It’s efficient for inserting and removing elements anywhere! 🚀
```go
ls := Linkedlist{}
ls.add(5)
ls.add(6)
ls.add(7)
ls.add(10)
ls.add(12)
ls.print()
ls.reverse()
ls.print()
```
Simple, efficient, and powerful! 💡

### 🔄 Doubly Linked List: Forward & Backward!
The Doubly Linked List is like the Singly Linked List's cooler cousin! Each node has two pointers: one to the next and one to the previous. You can traverse in both directions — like a true data ninja! 🥷
```bash
ls := Linkedlist{}
ls.add(5)
ls.add(6)
ls.add(7)
ls.add(10)
ls.add(12)
ls.print()
```
Now that’s some serious power right there! ⚡️
## 🔥 Data Structures Explained <a name="data-structures-explained"></a>
- __Stack__: A last-in, first-out data structure, ideal for problems like undo/redo operations or depth-first search (DFS).
- __Queue__: First-in, first-out, perfect for managing tasks in order, like handling requests in a web server or breadth-first search (BFS).
- __Singly Linked List__: A linear data structure with efficient insertions and deletions at the head or tail.
- __Doubly Linked List__: An extended version of the linked list where each node has pointers to both the next and previous nodes, enabling efficient two-way traversal.

## 🛠️ Contributing <a name="contributing"></a> 
We’d love to have you contribute! 💪 Whether it’s fixing a bug, adding a new feature, or improving documentation, every contribution is welcome! 😄
Let’s make this project even better together! 💥

## 📝 License <a name="license"></a> 

This project is licensed under the MIT License - see the [LICENSE]("https://opensource.org/license/mit") file for details. 🗝️
