Write some simple "Hello, world" code.

Use the go command to run your code.

Use the Go package discovery tool to find packages you can use in your own code.

Call functions of an external module.

---

`go.mod`: 
- when our code import packagegs from other modules, those dependencies are tracked via `go.mod` module which stays with our source code

`go mod init <module_path>`
- module path will typically be the repository location where your source code will be kept
- If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module


---

The Go Code:
- Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
- Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
- Implement a main function to print a message to the console. A main function executes by default when you run the main package.

```
a package is a directory of .go files, and packages form the basic building blocks of a Go program. Using packages, you organize your code into reusable units. 

A module, on the other hand, is a collection of Go packages, with dependencies and versioning built-in.
```

`go.sum` file for use in authenticating the module

When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version