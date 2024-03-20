# `berr`: Better Errors

Better errors replaces the standard error message with a more detailed one. 
It also adds a stack trace and a code snippet to the error message. 
This makes it easier to debug your code and understand what is going on.

Yes, message output is inspired by the rust's [`anyhow` crate](https://docs.rs/anyhow/latest/anyhow/). :D

## Installation

```shell
go get github.com/sebastianwebber/berr@latest
```

## Example

```shell
❯ go run examples/main.go
2024/03/20 12:09:31 INFO message example="simple error" error="simple error"
2024/03/20 12:09:31 INFO message example="complex error"
  error=
  │ complex error
  │
  │ caused by:
  │    0: simple error
2024/03/20 12:09:31 INFO message example="very complex error"
  error=
  │ very complex error
  │
  │ caused by:
  │    0: complex error
  │    1: simple error
2024/03/20 12:09:31 INFO message example="ultra complex error"
  error=
  │ ultra complex error
  │
  │ caused by:
  │    0: very complex error
  │    1: complex error
  │    2: simple error
2024/03/20 12:09:31 INFO message example="god like complex error"
  error=
  │ god like complex error
  │
  │ caused by:
  │    0: ultra complex error
  │    1: very complex error
  │    2: complex error
  │    3: simple error
2024/03/20 12:09:31 INFO message example="join error"
  error=
  │ simple error
  │
  │ complex error
  │
  │ caused by:
  │    0: simple error

```