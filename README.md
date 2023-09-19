# Slang programming language


- A simple programming language aimed to be used for educational purposes, and to teach how programming language interpreters
  can be made.

- Since it's made for educational purposes, all the parts that make a lexer are made from scratch `lexer`, `parser` etc.

## Language Example

- The language is a `C` derived language, with the ability to create `structs` as complex types, and declare functions as First class values.
- The language could be dynamically typed, to make it easier to implement, but we preferred the hard path: **IDE** 

```
// variable declarations
var a = 10;
var c = "some string";
var arrayValue = [1, 2, 3];


fn add(a:int, b:int) {
return a + b;
}

 func doSomething() {
        return a + c ;
    } 
}



var conditional = func(a, b) {
}
// a program entrypoint
func main() {
var c = a + add(1, 4);


// builtin functions
print(c);
[...]

}
```


## GRAMMAR

### Below are the grammar rules for the Slang.


