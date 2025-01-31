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

* To build an executable for your project, use -

    > go build

    It will find the main.go file, and build the executable

* To build for a particular OS, like "windows", "linux", "darwin"(MacOS) -

    > GOOS="desired_os" go build

* In linux, to run the executable, first, change its permission, using -

    > chmod +x filename

    then run the executable, using -

    > ./filename

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

## Memory Management 

* Go provides 2 methods, to allocate memory 
1. new() -
    
    * Allocates the memory but no Initialization
    * Zeroed Storage, you can't put any data intially
    * Gives a memory address

2. make() - 

    * Allocates the memory and initializes it
    * Non-Zeroed Storage, can put data an the time of initialization
    * Gives a memory address

## Pointers

* Syntax - Declare a pointer variable -

    > var <pointer_name> *<data_type>

    e.g.,
        
    > var ptr *int

* Initialize a pointer with a variable's address -

    > var <pointer_name> = &<variable_name>

    e.g.,

    > var ptr = &myNum

* To get the value stored at the location, where pointer 'ptr' is pointing -

    > *ptr

## Arrays

* Declare an array, Syntax -

    > var <array_name> [size_of_array]<data_type>

    e.g.,

    > var arr [4]string // Declares an array of length 4

* Initialize an array, with values, Syntax -

    > var <array_name> = [size_of_array]<data_type>{"val1", "val2", "val3"}

    e.g.,

    > var arr = [5]string{"C", "C++", "Java"}

* In Go, Arrays don't have any functionality like... sorting, etc.

## Slices

* Declare an empty Slice, Syntax -
    * Method 1-
        > var <slice_name> = []<data_type>{}

        e.g.,

        >	var fruitList = []string{}

        - Dont define the length/size of slice
	    - When using Method 1... you need to add {} and add some values if any...

    * Method 2-
        > var <slice_name> = make([]<data_type>, initial_size_of_slice)

        e.g.,

        > var vegList = make([]string, 4)

        - We can assign values to the slice, using indexes.. upto index 3 [0-3]
        - After that, we need to append the values, using the append function

* Indexing or Slicing in the slices is same as in Python
* Slices are Arrays with a layer of Abstraction and a bunch of additional functions 

* To Remove a value, at index, i, we use the append method, and the slicing concept, as -
    > arr = append(arr[:i], arr[i+1:]...)

* The `...` inside the parenthesis, denotes, the function will take multiple arguments