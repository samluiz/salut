# Salut

## Introduction

Salut is an interpreted programming language written in Go.

## Installation

WIP

## Getting Started

```
def x = 5;
def y = 2;

if (x - y == 3) {
  return x - y;
} else {
  return x + y;
}
```

## Syntax

### Types
WIP
### Data structures
WIP
### Variables:
For defining a variable in Salut, just use the keyword def:
  ```
  def x = 10;
  ```
### Operators:
  ```
  4 + 7 * 2 / 2 - 3;
  ```
### Conditional Statements:
Salut support if, else and will suport elif
  ```
  if (2 > 1) {
    ...
  } else {
    ...
  }
  ```
### Comparison:
Salut support all the common comparison operators:
  ```
  2 == 2;
  3 < 4;
  5 > 2
  3 >= 3
  2 <= 2
  ```
### Functions:
Salut supports first class functions, meaning that a function is just a value:
```
def add = f(x, y) {
  return x + y;
}

out(add(2, 3));
```
### Loops
WIP

## Examples

WIP

## License

WIP

## Contributing

WIP