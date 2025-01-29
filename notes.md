# Golang Notes
* **Compiled Language**
* Has a single entry point, `main`
* Has concept of Struct, in place of Classes
* Functions or variables, having *first Character in upper-case* are considered as *Public*, whereas, those with *first character in lower-case* are considered *Private*
* Everything is a *Type*

## Important Commands
* Before stating a project, initialize the module, using -  
 
    > go mod init <module_name>

    e.g.
    
    - go mod init hello
    - go mod init github.com/dhairyahans/hello   

* To Run the Go program -

    > go run main.go

* To get the path, where Go is installed -

    > go env GOPATH

## Variables

* Declare a variable, Syntax -

    > var `var_name` dataType


* Initialize a variable, Syntax -

    > var `var_name` dataType = `value`

    OR

    > `var_name` := `value`     // No var syntax

    The walrus operator style variable initialization, is allowed only inside functions in Local Context, and not in Global Context

* You can skip the dataType, but once defined, one can't change the dataType of the value, like- 

    > var name = "Dhairya"

    > name = 12     // Not allowed

* To declare a variable, that is public, i.e., Publically accessible to all the files, make the first character, upper case, as -

    > const LoginToken = "qwertyuio"


## comma,ok || comma,error SYNTAX - 
* Used in place of try-catch, as Go doesn't have the try-catch statements
* SYNTAX -

    > value, err := functionCall()

    e.g. -

    input, err := sum(2,4)

