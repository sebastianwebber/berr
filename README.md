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

![example](./examples.gif)
> check [`./examples/main.go`](./examples/main.go) for details.