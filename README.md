# Go-Language-Specialization

### Table of Contents

   * [Course 1 : Getting Started With Go](#getting-started-with-go)
     * [Composite Data Types](#composite-data-types-in-golang)
       * [Arrays](#array)
       * [Slices](#slices)
       * [Hash Tables](#hash-tables)
       * [Maps](#maps)
       * [Structs](#structs)
     * [Protocols and Formats](#protocols-and-formats)
   * [Course 2: Functions, Methods and Interfaces](#functions-methods-and-interfaces) 
     * [Funtions and Organizations](#funtions-and-organizations)
     * [Function Types](#function-types)
     * [Object Orientation in Go](#object-orientation-in-go)
     * [Interfaces for Abstraction](#interfaces-for-abstraction)
  


## Getting Started With GO

#### Composite Data Types in Golang

##### Array
  * Fixed length
  * can have only one type of data
  * indices start at 0
  * each array element is initialized to zero value in go

```go
 var arr [5]int
 arr[0] = 12

 // Array Literal: array with predefined values

	var y [5]int = [5]int{10, 20, 30} // Partial assignment
  x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
```

  * Iterating over arrays
  ```go
  for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
  }
  
  for index, element range(arr) {
    fmt.Println(index, element)
  }
  ``` 

##### Slices

 * Slice is a variable sized array.
 * *Pointer* indicates the start of the slice
 * *Length* of slice is the total number of elements in the slice
 * *Capacity* is the maximum number of elements in the slice, from start to end

  ```go
  cards := [...]int {1,2,3,4,5}
  card1 := cards[0:2]
  card2 := cards[2:]

  len(cards) // length of slice
  cap(cards) // capacity of slice
  // For example
  len(card1) // 2
  cap(card1) // 5 since it is assigned from cards
  ```

  Slice Example

  ```go

    package main

    import "fmt"

    func main() {
      cards := [...]int{1, 2, 3, 4, 5}
      fmt.Println(cards) // [1,2,3,4,5]
      card1 := cards[0:2] 
      fmt.Println(len(card1), cap(card1)) // 2,5
      fmt.Println(card1) // [1,2]
      cards[0] = 12
      fmt.Println(cards) // [12,2,3,4,5]
    }
  ```

  **`make()` to create slices and array**

  * 2 Argument Version : specify type and capacity
  ```go
    slice := make([]int, 10) // []int is slice of int and 10 is the capacity
    // length == capacity
  ```
  * 3 argument version: specify type, length and capacity 
  ```go
  slice := make([]int, 10, 15); // length is 10 and capacity is 15
  ```

  **`append()` to add elements to end of  slice**
  ```go
    slice = append(slice, 100); // append returns a new slice so do not forget to assign it to the slice we are adding elemets to

  ```

##### Hash Tables

  * Contains key:value pairs, where each key must be unique

##### Maps

  * Implementation of hash table
  * use `make()` to create a map
    ```go
    var idMap map[string]int  // key type is string and value type is int
    idMap = make(map[string]int)

    // or we can do it like
    idMap := make(map[string]int)
    ```
  * Initializing using a **map literal**
    ```go
    idMap := map[string]int {
      "salary" : 20000,
    }
    fmt.Println(idMap) // map[salary:20000]
    ```
  * Accessing the maps
  ```go
    fmt.Println(idMap["salary"])
  ```  
  * Adding/Updating a new key:value pair
  ```go
    idMap["account"] = 3210
    idMap["salary"] = 25000
  ```
  * delete a key:value pair using `delete()`
  ```go
    delete(idMap, "salary")
  ```
  * to check if a key exists or not
  ```go
    value, doesExist := idMap["salary"] 
    /*
     value is the value of salary key and deosExist is bool that's true if key exists 
     and false if does not exists and value will be zero in that case
    */
  ```
  * to find length of map
  ```go
    len(map)
  ```
  * Iterating over maps
  ```go
    key, value := range idMap {
      fmt.Println(key, value)
    }
  ```

##### Structs

  * Aggregate data type
  * Groups together other data types

  ```go
    type Person struct {
      name string // fields
      address string
    }

    var person1 Person
    fmt.Println(person1) // {}

    person1.name = "Hridayesh"
    x := person1.name 

  // Initializing a struct

  p1 := new(Person) // initialize fields to zero,p2 is a pointer to person struct

  p2 := Person{name: "Joe", address: "22 Baker's Street"}
  ```


#### Protocols and Formats

  1. RFC : Request for Comments
   * definition of internet protocols and formats 
   * for example 
     * HTML(RFC no 1866) etc
     * URI (3986)
     * HTTP (2616) 

  2. Golang provides specific packages to work with different RFCs for example net/http for http protocol
  3. JSON: JavaScript Object Notation, RFC 7159
     1. json is used for structured data representation and can be resembled to a map or struct in golang
        ```go

          type Person struct {
            name string
            addr string
          }

          p1 := Person(name: "joe", addr: "address")
         ```
      1. JSON is all unicode
      2. Human Readable and Fairly compact representation
      3. We can use golang structs and arrays to later convert them into JSON objects
      4. **JSON Marshalling** : converting an object into JSON 
        ```go
          import "encoding/json"
          p1 := Person(name: "Hridayesh", addre: "22 Baker's Street")
          byteArray, err := json.Marshal(p1)     
        ```
      5. **JSON UnMarshalling** : convert JSON into golang object
        ```go
          var p2 Person
          err := json.Unmarshal(byteArray, &p2)
        ```
  4. File Access : ioutil and os packages in golang


## Functions Methods and Interfaces

##### Funtions and Organizations

  1. Functions and Return Values
    ```go
    // return one value
      func foo(x int, y int) int {
        return x
      } 

    // returning multiple values

      func foo(x int, y int) (int, int) {
        return x, y
      }
    ```

  2. Call by Value and Call by Reference
     1. Call by Value
        1. Passed arguments are copied to parameters. 
        2. Arguments are copied to parameters and arguments are copies of originals and not original themselves

      2. Call by Reference
         1. Pass a pointer instead of passing the value itself
        ```go
          func main() {
            y := 12
            fmt.Println(y) // 12

            foo(&y)
            fmt.Println(y) // 14
          }

          func foo(x *int) {
            *x = *x + 2
          }
        ```
  3. Passing Array and Slices to Functions as Arguments
     1. Array Arguments are Copied. Can be a problem for larger arrays
     2. Passing array as Pointers 
        ```go
          import "fmt"

          func main() {
            y := [3]int{1, 2, 3}
            foo(&y)
          }

          func foo(x *[]int) {
            fmt.Println((*x)[0]) // notice the () around *x
          }
        ```
      3. Passing slices
         1. In golang slices are automatiicaly passed as pointer references. When you pass a slice it automatically copies the pointer to the slice.
         2. Slice is basically a pointer to an array with capacity and length.
         3. When passing slice there is no need to do any poniter operations
        ```go
          package main

          import "fmt"

          func main() {
            y := []int{1, 2, 3}
            foo(y)
          }

          func foo(x []int) {
            fmt.Println(x) // [1,2,3]
            x[0] = x[0] + 1
            fmt.Println(x) // 2,2,3
          }
        ```

#### Function Types
  1. **First Class Values**
     * Functions can be treated like any other types
     * Variables can be declared as function type
     * can be created dynamically
     * can be passed as argument and returned as values
     * can be stored in data structure
        1. Assigning func to a variable
      ```go
        import "fmt"
        func main() {
          inc := increment
          fmt.Println(inc(0)) // 1
        }

        func increment(x int) int {
          return x + 1
        }

      ```
        2. Declaring function as a variable
      ```go

        // declare a variable of type function
        var funcVar func(int) int
        // define a function
          func incFn(x int) int {
            return x+1
          }
        func main() {
          // assign the function to our function variable
          funcVar := incFn
          fmt.Pritln(funcVar(0))
        }

        // easier and better sntax

        func main() {
          // only works inside main function
          incFn :=  func(x int) int {
          return x + 1
        }

          inc := incFn
          fmt.Println(inc(0))
        }
      ```
        3. Functions as arguments
      ```go
       func funcArgs(function func(int) int) int {
        return function(10)
      }

      func main() {
        // passing anonymous function as argument
        r := funcArgs(func(x int) int {
          return x+2
        })
        
        fmt.Println(r) // 12
      }

      ```
##### Anonymous Functions: We do not need to name a function while passing it as an argument. We can just use anonymous functions like we used in `funcArgs`.

  2. **Returning Functions as Arguments**
      * Functions can return other functions as values
      ```go
        
      func increment(x int) int {
        return x+1
      }

      func returnIncrement() func(int) int {
        return increment
      }

      func main() {
        inc := returnIncrement()
        fmt.Println(inc(1)) // 2
      }
      ```

##### Environment of a Function: Go is lexically scoped
##### Closuer:  Go supports closure. A function and it's environment

  3. **Variadic and Deffered**
     1. Variable Number of Arguments
        1. Passing variable number of arguments to a function
        2. Using ... to signify variable no of arguments
        3. Treated as slice inside the function
        ```go
          func getMax(values ...int) int {
            values // [1,2,3,4,5]
          }

          getMax(1,2,3,4,5)
          // or use
          slice := []int{1,2,3,4,5}
          getMax(slice...)
        ```
      2. Deferred Function Calls
         1. function call can be deferred so that function is not executed when it get's called rather it is executed later on when surrounding function completes
         2. Typically used for cleanup activities
        ```go
          package main

          import "fmt"

          func multiply(x,y int) int {
            return x*y
          }

          func main(){
            // this will be executed in at the end of the main function
            defer fmt.Println("RESULT: ", multiply(1,2))
            fmt.Println("Hello Darkness")
          }
          // Hello Darkness
          // 2
        ```
        3. Arguments of deferred call are executed immediately
           ```go
            func main(){
              i := 0
              defer fmt.Println(i+1) // i+1 is already evaluated
              i++ // 3
              fmt.Println("Hello")
            }
            // Hello, 2 and not 3
           ``` 

#### Object Orientation in Go
##### Go is not an Object Oriented Language. There is no such thing as class in golang. It uses *type* to create custom types and *receiver functions* to associate functions to those types.
  1. Classes and Encapsulation
     1. **Classes** are basically data and methods combined together. 
        1. An object is basically an instance of the class that has some real data
     2. **Encapsulation** is hiding the data from the programmer. In general it's abstraction layer
        1. Data can be protected from the programmer
        2. Data can be only accessed using the methods exposed by the class
        3. The type and method both should be in the same package.
     ```go

        // create a custom type
        type Custom int
        // creating a receiver function for Custom type
        func (value Custom) Double() int {
          return int(value*2)
        }

        // custom struct type
        type Custom struct {
          fname string
          lname string
        }

        func (value Custom) Print() {
          fmt.Println(value.fname, value.lname)
        }

        func main() {
          // initialize Custom type variable
          v := Custom(3)
          // access method on Custom type using dot notation
          v.Double() // 6 
 
          var v Custom
          v.fname = "Hridayesh"
          v.lname = "Sharma"
          v.Print()
        }

      ```
  2. Encapsulation
     1. We can controll access to data in any of the packages by only exporting the public functions. For example in a package we can have a variable x and a function that manipulates the x. we can simlpy export the function.
     2. **Methods that start with a Capital letter are all public in a package and can be used outside.**
  3. Point Receivers
     1. In go data passed to methods is by value. That means it's copied to arguments.
     2. In case of **Point Receivers** we instead of object being passed, a reference to object is passed to the receiver function.
      ```go
        package main

        import "fmt"

        type Custom struct {
          fname string
          lname string
        }
        // instead of value now go will pass it as a reference to custom type
        func (value *Custom) Print() {
          // value here is a pointer
          fmt.Println((*value).fname, (*value).lname)
          // though it's a pointer we do not need to explictly dereference it 
          fmt.Println(p.fname, p.lname) // works correct as go automatically does that for us. Point is referenced as p and not *p
        }

        func main() {
          var v Custom
          v.fname = "Hridayesh"
          v.lname = "Sharma"
          // go will implicitly pass a reference to v to receiver function
          v.Print()
        }
      ```

#### Interfaces for Abstraction 


  1. Polymorphism in Go
     1. Ability of an object to be in different forms in different contexts. In OOP languages polymorphism is supported by inheritance. Each child class has a different implementation of function in parent class by something called function overriding. There is no concept of classes and inheritance in golang
  2. Interfaces
     1. Interface is a set of method signatures. i.e just declaration, no definitions
     2. It's just name, parameters and return values, no implementation
     3. Used to show conceptual similarity in types
     4. A Type satisfies an interface if type defines all the methods specified in the interface
      ```go
        // Shape2D is an interface
        type Shape2D interface {
          Area() float64
          Perimeter() float64
        }
        // type Circle satisfies the interface by implementing the methods of interface
        type Circle struct{
          radius float64
        }

        func (c *Circle) Area() float64 {
          return 3.14*c.radius*c.radius
        }

        func (c *Circle) Perimeter() float64 {
          return 2*3.14*c.radius
        }

        func main() {
          circle := new(Circle)
          circle.radius = 2
          fmt.Println(circle.Area())
        }

      ```
      3. Interfaces vs Concrete Types
         1. Concrete Types (builtin types: string, int, bool, float, flaot32, float64 etc)
            1. specify the exact represenation of data and methods. 
            2. complete method implementation is included
         2. Interfaces Types
            1. Specifies some method signatures, no implementation
            2. Can be treated like other values- assigned to variables, passed or returned from fns
            3. Interfaces Values have two components
               1. **Dynamic Type**:  Concrete type which it is assigned to
               2. **Dynamic Value** : Value of the dynamic type
                  1. For example when we create an interface Shape2D and then create a concrete type Circle which satisfies the Shape2D interface. Now when we assign a variable of type Circle to type Shape2D then the dynamic type is the Circle and dynamic value is the value of the Cricle variable.
                  ```go
                    package main

                    import (
                      "fmt"
                    )

                    type Speaker interface {
                      Speak()
                    }

                    type Dog struct {
                      name string
                    }

                    func (d Dog) Speak() {
                      fmt.Println(d.name)
                    }

                    func main() {
                        var s1 Speaker
                        d1 := Dog{name:"Brian"}
                        // s1 has dynamic type Dog and dyanmic value d1
                        s1 = d1
                        s1.Speak()
                        
                    }
                    // Dog type satisfies Speaker interface and therefore s1 can be assigned d1
                  ``` 
               3. Interface with nil dynamic value
                  ```go
                    var s1 Speaker
                    var d1 *Dog
                    s1 = d1 // d1 is nil here as it is a pointer and points to nothing for now as it has no data in it but it still has dynamic type Dog and we can do s1.Speak()
                  ```
               4. nil Interface Value
                  1. interface with nil dynamic type (interface that's not satisfied by any type)
                  2. very different from an interface with nil dynamic value
      4. Using Interfaces
         1. Ways to use interface
            1. Function which takes multiple types of parameters
            2. Function takes the interface as the type and different types of parameters must satisfy the interface
             ```go
              func FitInYard(s Shape2D) bool {
                if (s.Area() > 100 && s.Perimeter() >100) {
                  return true
                }
                return false
              }
              type Circle struct {
                radius float64
              }

              type Rectangel struct {
                height float64
                width float64
              }
              // Both Reactangle and Circle satisfy Shape2D
              FitInYard(rect) //
              FitInYard(circle) // 
             ```
         2. Empty Interface : Interface with no methods
      5. Type Assertions
         1. Interface hide the differences in types. For example FitInYard does not care if it's recctangle or circle
         2. Sometimes it can be required to expose the type on which interface is based on. We can a switch for that
            ```go
              func DrawShape(s Shape2D) bool {
                switch:= sh := s.(type) {
                  case Rectangle: 
                  case Circle
                  ...
                }
              }
            ```
      6. Handling Errors
         ```go
            type Error interface {
              Error() string
            }
          ```

      
