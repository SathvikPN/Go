In this tutorial you'll create two modules. The first is a library which is intended to be imported by other libraries or applications. The second is a caller application which will use the first.

This tutorial's sequence includes seven brief topics that each illustrate a different part of the language.
```
Create a module -- Write a small module with functions you can call from another module.
Call your code from another module -- Import and use your new module.
Return and handle an error -- Add simple error handling.
Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
Return greetings for multiple people -- Store key/value pairs in a map.
Add a test -- Use Go's built-in unit testing features to test your code.
Compile and install the application -- Compile and install your code locally.
```

---

In a module, you collect one or more related packages for a discrete and useful set of functions

Exported Names:  In Go, a function whose name starts with a capital letter can be called by a function not in the same package.

---

For production use, you’d publish the `example.com/greet` module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. 

For now, because you haven't published the module yet, you need to adapt(go mod edit) the `example.com/call` module so it can find the `example.com/greet` code on your local file system.

```Go
go mod edit -replace example.com/greet=../greetings
```

---

### Findings

Name of Directory -- filename -- package name
- identical by convention
- can be different. (see this directory code)

Structure:
A library: example.com/greet  (Module that provides functionality)
A application: example.com/call (Module of application which runs)

Directory name: greetings
filenames: `greetings.go` `greetings2.go` (files contain different code-parts belonging to same greet package)
package: greet 

So,
- package codebase can be splitted into many filenames.go 
- To import and use, modulename is the primary reference not filenames or dir names
- To use a locally existing module, mod file has to be edited by replace (replace lookup target from published-repo-check to local directory). One has to reflect the directory name here
