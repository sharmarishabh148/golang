#SOLID Principle


1. ##Single Responsibility Principle (SRP):##


A class should have only one reason to change, meaning it should have only one responsibility.
```
package main

import "fmt"

// Incorrect: Violates SRP
type Employee struct {
    name     string
    salary   float64
    isRetired bool
}

func (e *Employee) calculateSalary() {
    // Calculate salary
}

func (e *Employee) checkRetirementStatus() {
    // Check retirement status
}

// Correct: Follows SRP
type Employee struct {
    name     string
    salary   float64
    isRetired bool
}

func (e *Employee) calculateSalary() {
    // Calculate salary
}

type HR struct{}

func (hr *HR) checkRetirementStatus(e *Employee) {
    // Check retirement status
}

```
2.##Open/Closed Principle (OCP):##


Software entities (classes, modules, functions) should be open for extension
but closed for modification.

```
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func CalculateArea(s Shape) float64 {
    return s.Area()
}
```
You can add new shapes that implement the Shape interface without
modifying the CalculateArea function.
3.##Liskov Substitution Principle (LSP):##


Subtypes must be substitutable for their base types without altering
the correctness of the program.
In Go, this principle is closely related to interfaces and type assertions.
```
package main

import "fmt"

type Bird interface {
    Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
    fmt.Println("Sparrow can fly")
}

type Penguin struct{}

// Uncommenting this function would violate LSP.
// func (p Penguin) Fly() {
//     fmt.Println("Penguin can't fly")
// }

func LetItFly(b Bird) {
    b.Fly()
}
```
The Penguin type doesn't implement the Fly method because penguins can't fly.
This is an example of adhering to the LSP; it doesn't force subtypes to do 
something they can't do.

4.##Interface Segregation Principle (ISP):##


Clients should not be forced to depend on interfaces they do not use. This
principle is not strictly applicable in Go since Go uses implicit interfaces,
and types implement only the methods they need.

However, you can still design your interfaces to follow a similar idea. 
For example, if you have a large interface, you can break it down into smaller
interfaces that are more specific to each client's needs.

5.##Dependency Inversion Principle(DIP):##


A. High-level modules should not depend on low-level modules. Both should depend
on abstractions.
B. Abstractions should not depend on details. Details should depend on 
abstractions.
```
package main

import "fmt"

type MessageSender interface {
    SendMessage(message string)
}

type EmailSender struct{}

func (e EmailSender) SendMessage(message string) {
    fmt.Println("Sending email:", message)
}

type SMS struct {
    Sender MessageSender
}

func (s SMS) SendSMS(message string) {
    s.Sender.SendMessage(message)
}

```