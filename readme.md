A not-so-new <i>Go</i> Application <b>arises</b>.

# First Considerations

I'm using as references:
- [x] Alura Classes (Go Course),
- [x] Go Programming for Dummies by **Wei-Meng Lee**.
- [x] The lovely Go Documentation by **Google Team**.

___

# Getting Started

'OOP stands for Object-Oriented Programming, not Class-Oriented Programming.'

## What is Golang

Go was designed by three engineers at Google who were frustrated with the various toolsets that they were working on and set out to design a new language that would address the criticism of other languages while at the same time keeping their useful features.

Go was designed to ->
- [x] Use *static typing* and have the runtime efficiency of C.
- [x] Have the *readability* and *usability* of languages like Python and JavaScript.
- [x] Exhibit great performance in *networking* and *multiprocessing*.
- [x] Have fewer features that there's *only one* right way to solve a problem.

___

# Go File Structure

All files in the same directory belong to the same package, that's something to note. Despite which file you're coding on, as long as it's the same package as the other one you're referencing (via a function or a variable), you don't need to import the file. But if it's a different package, then you'll have to import and use the call by using exportedPackageName.functionName(), etc.

___

# Instantiating a Module

Run `go mod init <project_name>` inside the folder your main.go is at. Important to note that project_name is usually github/<your_git_username>/<repository_name>
or something like the website that you'll use to host your server. Main rule is
try to be as unique as possible.

When a go.mod file is in the root of your project directory, your project becomes
a *module*, and this brings advantages:
1. It makes it much easier to manage third-party dependencies.
2. Avoid supply-chain attacks.
3. Ensure reproducible builds of your application in the future.

___

# Building

`go build main.go` builds the project, creating a `main.exe`.  
`go run main.go` builds and runs the `main.exe` right after.

Querying command `go env` inside your project file, you will se some attributes to
the environment Go is setting you on. To compile for other Operating Systems, you'll have to declare `GOOS` AND `GOARCH` when using the command `go build`.

Windows: `GOOS=windows GOARCH=amd64 go build -o project.exe`.  
Linux: `GOOS=linux GOARCH=amd64 go build -o project`.
Apple: `GOOS=darwin GOARCH=amd64 go build -o project`.  

___

# Variables

In Go, regarding variables, a compilation error occurs when a:
1. a variable is declared but not used;
2. a variable is declared more than one;
3. a variable is assigned with a new value of a different type previously declared;
4. a variable is assigned with a value the type isn't capable of holding.
5. a *const* variable is assigned a new value.


- `var var_name type` and `var var_name type = var_value` are both *explicit typing*.
- `var var_name = var_value` is a *type inferred*.
- `var_name := var_value` is *short variable declaration*, and can't be used outside of functions.


## Nil-Safety

Given that a variable can be initialised but does not need to be assigned a value, then the value is always set as a default. For number-types, value 0 is default, for *string* it's " ", for *bool* it's false.

## _

Use `_` as a variable-name when you do not plan on using that variable.

## Type Conversions

As soon as you use a built-in function, import line is updated, so as long as you use any of these, the corresponding package will be imported automatically.

`strconv.aoti` converts string to int, `strconv.itoa` goes the other way.

Note that `strconv.aoti` returns two values, one value for the result and one for a possible error. Same happens to `strconv.parseBool` that converts to Bool, `strconv.ParseInt` and `strconv.ParseFloat`.

___

# Control Flow

## If Else

It's just the same as any other language. Only two differences are:

1. There's no ternary operator, and it's been stated that the reason for that is for devs to avoid creating so many ifs (mostly complex ones) just because they're shorter.
2. You must not encapsulate the condition inside (), but curly brackets are still to be used to write instructions in.

An improvement that is forgotten in other languages is *short-circuiting* in Go evaluations. This happens, for example, inside an if statement of two conditions with a logical OR operator `||`, in case the first condition returns true, the second condition is not evaluated because the outcome will not change. The if statement will also stop midway if the logical operator is AND `&&` and the first condition is false, which renders useless to evaluate the second condition given the result will definitely not change.

## Switch

When you have too many conditions to evaluate, a switch statement is recommended. Although it does seem the same as other switch statements in alike programming languages, a break is not needed in Go.

## For

Go is pretty not stupid, and that's why instead of having more than 27 different keywords for looping, Go only has one: *for*.

The syntax is quite like C++, except in a more ?gonic? way, we tend to use short variable declaration, so usually it'll be:

```go
for i := 0; i < 10; i++ 
```

And yes, no need for () to encapsulate *for* parameters either. In fact, if you encapsulate it, an error will occur. A word about increment/decrement operators is that it's only valid to use the sign AFTER the variable. In case you need an infinite loop, just use:

```go
for { 
    // instructions
}
```

In case you need to iterate through an iterable object you can use the keyword range:
```go
for index, value := range random_map_slice_array {}
```

You can also break inside a *for* loop. Beware that a break inside a inner loop will only break out of the inner loop, if there's to break completely, best way to  achieve this is by labeling the outer loop and using the keyword to break the labeled loop.

``` go
Outerloop: for i := 1; i < 10; i++ {
    for j := 1; j < 10; j++ {
        break Outerloop
    }
}
```

___

# Functions

The keyword for functions is func, it can return from zero to multiple values, and can have from zero to many parameters.

Unlike other languages, Go does not support *function overloading* in an attempt to keep Go simple.

There's a way of using optional parameters and that's with *variadic parameter*. It should be always be the last parameter though, otherwise an error will occur. You can either opt out of using that parameter or using it many times. A calculator function is a good example for *variadic parameter*. 

___

# Data Structures

As the name (and formation) suggest(s), a *Data Structure* is a specialised format for organising, processing, retrieving and storing data.  
As for Go, *arrays* are best for sequential data, *slices* for dynamic data, *maps* for key-value data, and *structs* for complex data.

## Array and Slices

An *array* is a numbered sequence of a specific length of items of the same type. The values an *array* holds are known as items or elements, and after an *array* is created, it cannot change in size.
Declaration of an *array* demans size:

```go
array [5]string = {"a", "r", "r", "a", "y"}
```

A *slice* is exactly an *array*, but it can change its size.
Declaration of a *slice* must not start with size specification.

```go
fib []int = {1, 2, 3, 5, 8}
```

Technically, a *slice* is a view into an underlying *array*, and it's represented by a *slice header*, which contains three fields:
- ptr: A *pointer* that points to the address of the underlying *array*.
- len: The *length* of the *slice*.
- cap: The *capacity*, maximum number of allowed elements in the *array*.

You can use `len(fib)` and `cap(fib)` to retrieve these values. Right after a *slice* is initialised, the `len()` and `cap()` will be the same. In the case of `fib`, `len()` and `cap()` will both be 5.

To append to a *slice*, do as following:

```go
fib = append(fib, 13, 21)
```

Note that you can append more than one value. But what happens to the `len()` and `cap()`?
Not surprisingly, `len()` has gone from 5 to 7 given that two more values were added to the *slice*, but cap didn't go along the same route.
Value of `cap()` went from 5 to 10, because what happened was that when a *slice* has reached its capacity, appending more items to it causes the *slice* to point to a new underlying *array*. All the existing items in the *slice* are copied onto the new *array*, together with the newly added items, therefore making the new `cap()` double what it was before the append.

## A bit about pointers

If you assign an *array* to another *array*, you create an independent *array* and neither *array* will share underwent changes to the other one, because all that was done was to create a copy.
If you do the same to a *slice* though, you'll create a destination *slice* that points to the same memory address of the source *slice*, therefore making it look more like an alias than anything else.

You could also just use the following for arrays and other variables:
```go 
copy(destination, source) 
``` 
or you could use * (*pointers*) and & (*references*).

## Struct

A *struct* is a data structure that acts just like a class, so no much more explaining has to be done then.

Some key-points tho:
1. Unlike other languages, Go has no *default* constructor.
2. Just like a *slice*, if you assign a *struct object* to another *struct object*, they both point at the same *struct object*. Any changes in one will reflect on both.
    - Note that passing a *struct* as a func parameter, changes inside the func will not reflect. For that to happen, use
    ```go
    type point struct {
        x float32
        y float32
    }

    func newPoint(x, y, z float32) *point { // this works as constructor
        p := point{x: x, y: y}
        return &p
    }

    func (p *point) move(deltax, deltay, deltaz float32) { // this works as a method
        p.x += deltax
        p.y += deltay
    }

    func main() {
        p := newPoint(7.8, 9.1)
        fmt.Println(p.length())
        p.move(0.1, 0.1)
        fmt.Println(*p) // {7.9 9.200001}
    }
    ```
    *code snippet plagiarised from Chapter 7 of Go Programming Language for Dummies.*

## Maps

As for an *array* or a *slice*, the position of items is important, and items are accessed by their locations. A *map*, on the other hand, is a hash table that stores data in an associative manner. Items in a *map* are not accessed according to their positions. Instead, you use _keys_, a set of unique value that identifies the elements in a _map_. Adding and removing item is extremely easy in this data structure.


Maps are easy but lengthy and I'm not sure I want to through all of it now. I might add more on this readme.md in the future, who knows.