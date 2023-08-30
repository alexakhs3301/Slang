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


fn add(a, b) {
return a + b;
}

fn isBigger(a, b) {
return a > b;
}

struct Data {
fn init(a, b, c) {
this.a = a;
this.c = a + b;
}

    func doSomething() {
        return a + c ;
    } 
}
// STRUCT_DECL: struct Identifier BLOCK
// BLOCK: { functions* }

//
// var data = new Data(a, b, c);

// data.doSomething();

// Object declaration
var object = new();
object.field = "field value";

object.method = func(a) {
return a * 2;
}

var value = object.call();

var loopConstruct = func(a, b) {
var sum = 0;
for i = 1; i < 10; i += 1 {
sum += i;
}
return sum;
}

var conditional = func(a, b) {
}
// a program entrypoint
func main() {
var c = a + add(1, 4);
// builtin functions
println(c);
}
```


## GRAMMAR

### Below are the grammar rules for the Slang.

// Parser rules
1. [ ] TODO

