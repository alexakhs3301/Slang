# SLANG

Slang supports functions, various datatypes, operators and statements.

## Functions
All function should be above of main code. You cannot call function if exists after calling.

```
fn add(x:int, y:int) : int {
    return x + y;
}
```

In this version we do not support accessability modifiers (public, private etc);

## Operators

### Arithmetic Operators

| Operator | Name |Description | Example
| -------- | ----- | ------ | ----- |
| `+` | Addition | Adds together two values | x + y |
| `-` | Substracts | Subtracts one value from another | x - y |
| `*` | Multiplication | Multiplies two values | x * y |
| `/` | Division | Divides one value by another | x / y |
| `%` | Modulus | Returns the division remainder | x % y |
| `^` | Power | Returns the power | x ^ 2 |


### Comparison Operators
| Operator | Name |Description | Example
| -------- | ----- | ------ | ----- |
| `==` | Equal to | Checks equality of two values | x == y |
| `!=` | Not Equal | Checks if two values are not same | x != y |
| `>` | Greater than | Checks if the left value is greater than right value | x > y |
| `<` | Less than | Checks if the left value is less than right value | x < y |

### Logical Operators
| Operator | Name |Description | Example
| -------- | ----- | ------ | ----- |
| `!` | Logical not | Reverse the result, returns lie if the result is truth | !truth |


### Bitwise Operatos
| Operator | Name | Description | Example
| ---- | ---- | ---- | ---- |
| `&` | Bitwise AND | inary AND Operator copies a bit to the result if it exists in both operands. | Input: 10 & 4 Expected: 0 |
| `\|` | Bitwise OR | Binary OR Operator copies a bit if it exists in either operand. | Input: 1034 \| 4 Expected: 1038 |
| `#` | Bitwise XOR | Binary XOR Operator copies the bit if it is set in one operand but not both. | Input: 10 # 4 Expected: 14 |
| `~` | Bitwise NOT | Binary Ones Complement Operator is unary and has the effect of 'flipping' bits. | Input: ~10 Expected: -11 | 

## Data Types

| Type | Name | Values |
| ---- | ---- | ---- |
| `int` | Integer| -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `bool` | Boolean | truth, lie |
| `string` | String | "text" |
| `[ ]` | Arrays | [1,2,3,4] |

### Strings
---
| Function | Description |
| ---- | ----|
| `first(string)` | Returns the first letter of string |
| `string + string` | Returns concatenated string |


### Arrays
---
| Functions | Description |
| ---- | ----|
| `first(array)` | Returns the first element of the array |
| `last(array)` | Returns the last element of the array |
| `rest(array)` | Returns all values after first element of the array |
| `push(array, newValue)` | Pushes new value in the last position of the array |
| `array[index]` | Gets the value of the called index |

## Includes

You are able to include files with functions

```
#add "filePath/fileName"
```

For example if you want to include a file which is under `extra` directory and the filename is `functions.slang` you need write this:

```
# add "extra/functions"
```

and you have include the file.

## How to print

In Slang there are two ways to print. You can use the built-in function `printer(object)` or you can just write the name of your variable e.g. `var x = 5; x;`. Both approches are correct.

## Comments

We support only line comments.
```
// This is a comment
```

# How to write Slang
In Slang you are not obliged to start with `main()` function. For example:
```
var x = 5;
x;
// It will print 5.
```
or
```
fn main() {
    var x = 5;
    x;
}
```

Both approches are correct.

# Your first program in Slang

```
fn add(x:int, y:int) :int {
	return x + y; 
}

fn pow(x:int) {
    return x^2;
}

var array = [1, 10, 25, 54];

if (pow(10) > 20) {
    add(last(array), first(array));
}

// it will print 55
```
