# Quick Start to Fract

+ [Keywords](https://github.com/fract-lang/fract/blob/master/docs/Fract/keywords.md)
+ [Data Types](https://github.com/fract-lang/fract/blob/master/docs/Fract/data_types.md)
+ [Operators](https://github.com/fract-lang/fract/blob/master/docs/Fract/operators.md)
+ [Embed Functions](https://github.com/fract-lang/fract/tree/master/docs/fract/embed%20functions)
+ [Imports](https://github.com/fract-lang/fract/blob/master/docs/fract/import.md)

## Comments
``#`` Is used for singline comments.

> If you want to show it as a comment, a space or similar must be after ``#`` otherwise it will be accepted as a macro.

### Examples
```
# Comment <- Singline comment.
#Comment  <- Macro
```
```
555 + 5 # Sum.
```
``#> ... <#`` are used for multiline comments.
### Examples
```
#>
  Hello Function
  Desc: Print hello to screen.
  func Hello()
    ...
<#
```

### Important
It is always advantageous to use a single line comment because the lexer adds free space to the memory by deleting these comments from the lines that are read.

## Process Priority
Fract adheres to transaction priority!
### Examples
```
5 + 2 * 2     # 9
(5 + 2) * 2   # 14
```

## Print "Hello World!"

[print function](https://github.com/fract-lang/fract/blob/master/docs/fract/embed%20functions/print.md)

### Examples
```
print(5555 + 1) # Print 5556
```
```
var x int32 = 5
print(x) # Print 5
```

## Statement Terminator
With the Statement terminator, you can perform multiple operations on the same line without moving to a new line.

### Syntax
```
[STATEMENT]; [STATEMENT]; [STATEMENT];...
```

### Examples
```
print(5); print(2)   # Print 5 and 2
```

## Range Decomposition
Until the brackets are closed, they are tokenized.

> For example, if you have endless long conditions in conditional statements, you can use parentheses to use the bottom lines!

### Examples
```
print(4 +
      4)        # Tokenizer Result: print(4+4)
```

## Variables
### Definition

> A value must be given when defining a variable!

> Variable names are must comply to [naming conventions](https://github.com/fract-lang/fract/blob/master/docs/Fract/naming_conventions.md).

> cannot change values of const variables!

### Syntax
```
var [NAME] = [VALUE]
```
```
const [NAME] = [VALUE]
```
### Examples
```
const Pi = 3.14
```
```
var (
  FibonacciFirst = 1,
  Ln2Hi          = 6.93147180369123816490e-01,
  Hex            = 0x07EDD5E59A4E28C2
)
```

## Set Defined
### Syntax
```
[NAME] = [VALUE]
```
### Examples
```
var a = 45      # Value is 45
a = 1           # Value is 1

var (
  c = 3,        # Value is 3
  d = 5,        # Value is 5
  e = 9         # Value is 9
)
```

## Deletion Defines
You can free space and customize usage by deleting definitions from memory.

> Fract does not allow direct memory management! You can contribute to usage by deleting only memorized definitions.

### Syntax
```
del [NAME], [NAME], [NAME],...
```
### Examples
```
var a = 4
var a = 5        # Error, defined name 'a'

------------------

var a = 4
del a            # Remove 'a' variable from memory
var a = 5        # No error, a is 5
```
```
var (
  a = 0,
  b = 0,
)
del a, b         # Remove 'a' and 'b' variables from memory
                 # No defined variables

------------------

# Function removing

del a()          # Remove 'a' function from memory
del a, a()       # Remove 'a' variable and function from memory
```

### Protected Keyword
Protected objects is cannot remove manually from memory.

#### Examples
```
protected var example = 4
del var # Error: Protected objects cannot be deleted manually from memory!
```

## Arrays
They are structures that can hold more than one value in arrays. An array has a limited size and this size is determined at creation time. <br>
Syntax for creating an array that characterizes with 4, 5, 6, 7 elements:
```
var array = [ 4, 5, 6, 7 ]  # Elements: 4, 5, 6, 7
```
The syntax for accessing an element of an array with index:
```
array[index]
```
The syntax for setting an element of an array with index:
```
array[index] = value
```

### How can you quickly use the data in an array for arithmetic operations?
```
var array = [ 0, 4, 4, 2 ]          # Elements: 0 4 4 2
array += 5                          # Elements: 5 9 9 7
```
```
var array = [ 0, 4, 4, 2 ]          # Elements: 0 4 4 2
var array2 = [ 2, 2, 2, 2 ]         # Elements: 2 2 2 2
array = array + array2              # Elements: 2 6 6 4
```

> An array can be manipulated with an arithmetic value. However, when executing with a different array, the array must have only one element or the same number of elements.

## Conditional Expressions
You can let the algorithm flow with the conditions. Fract offers the If-Else If-Else structure like most programming languages.
"If" is the main condition, the alternative conditions that will come later are shown as "Else If".
When one condition is fulfilled, other conditions do not. Therefore, "If" must be rewritten each time to create a different order of conditions.

> All kinds of data can be given, but conditioning only looks for 0 and 1. 0 is accepted as false, 1 true.

### Syntax
```
if [CONDITION]
end
```
```
if [CONDITION]
  # ...
elif [CONDITION]
  # ...
elif [CONDITION]
  # ...
else
 # ...
end
```

A condition can be given any kind of value, but it only works with true(1) and false(0).
Unlike most languages, you won't get an error even if you only enter an integer value in the condition. It looks at the value and if it is 1 it fulfills the condition.
```
var example = 0
if example
  # ...
end
```

## Loops

Repetitive operations can be done using loops.

### While Loop
The while is a loop that happens as long as the condition is met.

#### Syntax
```
for [CONDITION]
  # ...
end
```
#### Examples
```
var counter = 0
for counter <= 10
  print(counter)
  counter = counter + 1
end
```

### Foreach Loop
You can rotate the elements of arrays one by one with the foreach loop.

#### Syntax
```
for [VARIABLE_NAME_INDEX] in [VALUE]
  # ...
end
```
```
for [VARIABLE_NAME_INDEX], [VARIABLE_NAME_ELEMENT] in [VALUE]
  # ...
end
```
#### Examples
```
var t1 = [ 0, 3, 2, 1, 90 ]
for index in t1
  print(t1[index])
end
```
```
var t1 = [ 0, 3, 2, 1, 90 ]
for index, item in t1
  print(index, fin=" | ")
  print(item)
end

# OUTPUT
0 | 0
1 | 3
2 | 2
3 | 1
4 | 90
```

### Tip
The best way give an empty condition for infinite loop.
#### Examples
```
for
  # ...
end
```

### Break Keyword
With the keyword break, it is possible to override and terminate the entire loop even when the loop could still return.
#### Examples
```
var counter = 0
for counter <= 10
  counter = counter + 1
  if counter > 5
    break
  end
  print(counter, fin=" ")
end

# Output: 0 1 2 3 4 5
```

### Continue Keyword
It can be used to pass the cycle to the next cycle step. If there is no next loop step, the loop is terminated.
```
for _, item in [ 0, 1, 2, 3, 4.0 ]
  if item == 1 || item == 3
    continue
  end
  print(item, fin=" ")
end

# Output: 0, 2, 4
```

### Loop and else blocks
Highly functional loop else blocks can be defined arbitrarily. <br>
These blocks execute when the array is empty for foreach loops and the loop never returns in while loops.

#### Examples Syntaxes
```
for bool
  # ...
else
  # ...
end
```
```
for _ in array
  # ...
else
  # ...
end
```

## Functions
Functions are very useful for adding functionality to your code.

### Syntax
Define:
```
func [NAME]([PARAM], [PARAM], [PARAM],...)
  ...
end
```
Define with default values:
```
func [NAME]([PARAM], [PARAM]=[VALUE],[PARAM]=[VALUE],...)
  ...
end
```
Multi valued parameters:
```
func [NAME](...[PARAM])
  ...
end
```
Call:
```
[NAME]([PARAM], [PARAM],...)
```
Call with parameter setter:
```
[NAME]([PARAM_NAME]=[VALUE], [PARAM_NAME]=[VALUE],...)
```

### Examples
```
func multiValuedParams(...values, s)
  print(values + s)
end

multiValuedParams(4,4,4,4,4,4, s=3)
```
```
func prime(x)
  if x == 2
    ret true
  end

  for _, y in range(2, x-1)
    if x % y == 0
      ret false
    end
  end
  ret true
end

for _, number in range(0, 10)
  print(number, fin=" ")
  print(prime(number))
end
```
```
func printHello()
  print("Hello")
end

printHello()
```

### Ret Keyword
The keyword ret is used to return the value of the function.

#### Syntax
```
ret [VALUE]
```

#### Examples
```
func reverse(x)
  ret x * -1
end

reverse(-500) # Returns: 500
```

## Exception Handling
### Syntax
```
try
  # ...
end
```
```
try
  # ...
catch
  # ...
end
```


## Macros

### Macro Conditions

+ [Macro Condition Variables](https://github.com/fract-lang/fract/blob/master/docs/fract/macro_if_variables.md)

#### Examples
```
#if OS == "windows"
  const os = "Windows"
#elif OS == "darwin"
  const os = "Darwin"
#else
  print("ERROR: Invalid operating system!")
#end

print(os)
```

### Macro Defines
You can define variables for macro conditions.

> Default value is ``false``.
 
> Macro defines is not delete like variables or functions if defined in code block or like.

#### Examples
```
#define PRINT_AS_FLOAT true
#define PRINT_AS_OBJECTSTR

const value = 4.535

#if PRINT_AS_FLOAT
  #if PRINT_AS_OBJECTSTR
    print(string(value, type="object"))
  #else
    print(value)
  #end
#else
  #if PRINT_AS_OBJECTSTR
    print(string(int(value), type="object"))
  #else
    print(int(value))
  #end
#end
```
