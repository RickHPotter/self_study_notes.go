A not-so-new <i>Go</i> Application <b>arises</b>.

# First Considerations

I'm using as references:
- [x] Alura Classes (Go Course),
- [x] Go Programming for Dummies by **Wei-Meng Lee**.
- [x] The lovely Go Documentation by **Google Team**.

___

# Getting Started

'OOP stands for Object-Oriented Programmin, not Class-Oriented Programming.'

## What is Golang

Go was designed by three engineers at Google who were frustrated with the various toolsets that they were working on and set out to design a new language that would address the criticism of other languages while at the same time keeping their useful features.

Go was designed to ->
- [x] Use *static typing* and have the runtime efficiency of C.
- [x] Have the *readability* and *usability* of languages like Python and JavaScript.
- [x] Exhibit great performance in *networking* and *multiprocessing*.
- [x] Have fewer featuress that there's *only one* right way to solve a problem.

___

# Go File Structure

All files in the same directory belong to the same package, that's something to note. Despite which file you're coding on, as long as it's the same package as the other one you're referencing (via a function or a variable), you don't need to import the file. But if it's a different package, then you'll have to import and use the call by using exportedPackageName.functionName(), etc.

___

# Instantiating a Module

Run `go mod init <project_name>` inside the folder your main.go is at. Important to note that project_name is usually github/<your_git_username>/<repository_name>
or something like the website that you'll use to host your server. Main rule is 
try to be as unique as possble.

When a go.mod file is in the root of your project directory, your project becomes
a *module*, and this brings advantages:
1. It makes it much easier to manage third-party dependencies.
2. Avoid suply-chain attacks.
3. Ensure reproducible builds of your application in the future.

___

# Building

`go build main.go` builds the project, creating a `main.exe`.  
`go run main.go` builds and runs the `main.exe` right after.

Querying command `go env` inside your project file, you will se some attributes to
the environment Go is setting you on. To compile for other Operating Systems, you'll have to declare `GOOS` AND `GOARCH` when using the command `go build`.

Apple: `GOOS=darwin 	GOARCH=amd64 go build -o project`.  
Wndws: `GOOS=windows 	GOARCH=amd64 go build -o project.exe`.  
Linux: `GOOS=linux 		GOARCH=amd64 go build -o project`.

___

# Variables

In Go, regarding varibles, a compilation error occurs when a:
1. a variable is declared but not used;
2. a variable is declared more than one;
3. a variable is assigned with a new value of a different type previously declared;
4. a variable is assigned with a value the type isn't capable of holding.
5. a const variable is assigned a new value.


- `var var_name type` and `var var_name type = var_value` are both *explicit typing*.
- `var var_name = var_value` is a *type infered*.
- `var_name := var_value` is *short variable declaration*, and can't be used outside of functions.


## Nil-Safety

Given that a variable can be initialised but does not need to be assigned a value, then the value is always set as a default. For number-types, value 0 is default, for strings, it's " ", for boolean is false.

## _

Use _ as a variable-name when you do not plan on using that variable.

## Type Conversions

As soon as you use a built-in function, import line is updated, so as long as you use any of these, the corresponding package will be imported automatically.

`strconv.aoti` converts string to int, strconv.itoa goes the other way.

Note that `strconv.aoti` returns two values, one value for the result and one for a possible error. Same happens to `strconv.parseBool` that converts to Bool, `strconv.ParseInt` and `strconv.ParseFloat`.

___

# Control Flow

## If Else

It's just the same as any other language. Only two differences are:

1. There's no ternary operator, and it's been stated that the reason for that is for devs to avoid creating so many ifs (mostly complex ones) just because they're shorter.
2. You must not encapsulate the condition inside (), but curly brackets are still to be used to write instructions in.

An improvement that is forgotten in other languages is *short-circuiting* in Go evalutations. This happens, for example, inside an if statement of two conditions with a logical OR operator (||), in case the first condition returns true, the second condition is not evaluated because the outcome will not change. The if statement will also stop midway if the logical operator is AND (&&) and the first condition is false, which renders useless to evaluate the second condition given the result will definitely not change.

## Switch

When you have too many conditions to evaluate, a switch statement is recommended. Although it does seem the same as other switch statemenets in alike programming languages, a break is not needed in Go.

## For

Go is pretty not stupid, and that's why instead of having more than 27 different keywords for looping, Go only has one: *for*.

The syntax is quite like C++, excecpt in a more ?gonic? way, we tend to use short variable declaration, so usually it'll be:

```go
for i := 0; i < 10; i++ 
```

And yes, no need for () to encapsulate for parameters either. In fact, if you encapsulate it, an error will occur. A word about increment/decrement operators is that it's only valid to use the sign AFTER the variable. In case you need an infinite loop, just use:

```go
for { 
    // instructions
}
```

In case you need to iterate through an iterable object you can use the keyword range:
```go
for index, value := range random_map_slice_array {}
```

You can also break inside a for loop. Beware that a break inside a inner loop will only break out of the inner loop, if there's to break completely, best way to  achieve this is by labeling the outer loop and using the keyword to break the labeled loop.

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