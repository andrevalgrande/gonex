# gonex
Bizarre as it seems, gonex is an exception handling package for the Go language.

## Install

```
go get -u github.com/andrevalgrande/gonex
```

## Basics
Gonex violates basic go language philosophy in order to provide rudimentary exception handling. That's mostly because I got tired of the whole "return error / print error / etc" routine. So let's get down to business!

## Throwing

```go
import "github.com/andrevalgrande/gonex"

func main() {
	exceptions.Throw("OMG something went wrong!")
}
```

This will cause your application to panic, because you've thrown an exception but didn't do any handling.

## Handling
Say you have a Person struct and you want to make sure that the first and last names are not blank.
```go
package main

import "github.com/andrevalgrande/gonex"

type Person struct{
	FirstName string
	LastName string
}

func Validate(person Person){
	if (person.FirstName == "") || (person.LastName == ""){
		exceptions.Throw("You should inform first and last name.")
	}
}

func main() {
  defer exceptions.Catch(func(){
    fmt.Println("Exception catched!")
    fmt.Println(exceptions.GetString(true))
  })
	person := Person{}
	Validate(person)
}

```
Note that the exceptions.Catch() should always be used as a defer. It's also a good idea to plan your catches early in the method's code.

## Append warning
A more refined approach to the previous example:
```go
package main

import "github.com/andrevalgrande/gonex"

type Person struct{
	FirstName string
	LastName string
}

func Validate(person Person) (isValid bool){
  isValid = ture
	if person.FirstName == ""{
		exceptions.AppendWarning("You should inform first name.")
		isValid = false
	}
	
	if person.LastName == ""{
		exceptions.AppendWarning("You should inform last name.")
		isValid = false
	}
	
	return
}

func main() {
  defer exceptions.Catch(func(){
    fmt.Println("Exception catched!")
    fmt.Println(exceptions.GetString(true))
  })
  
	person := Person{}
	if!Validate(person){
	  exceptions.Throw("The informed person is not valid.")
	}
}

```
This way you can validate for multiple scenarios before actually raising an exception. This approach also allows for nested code to append its own warnings, so that when the parent call throws an exception, all warnings will be thrown with it.

##Getting information from the exception
In the previous examples, we have a call to exceptions.GetString(). This will return a string with all the exception data. There is a single boolean parameter, includeStack, which appends the call stack at the moment of the throw.

##Clearing exception data
Simply call exceptions.Clear()
