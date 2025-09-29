# Creating Custom Roach Builds

The Roach interpreter is written in Go, which gives you the powerful ability to extend its functionality by binding functions from any Go package. While the standard `roach` binary comes with a rich set of built-in functions, you can create a custom build to add your own or to pull in features from the vast Go ecosystem.

This guide will walk you through the process of creating a custom Roach interpreter with your own Go functions baked in.

## Prerequisites

- A working Go development environment (Go 1.18 or later).
- Basic familiarity with Go syntax.

---

## Step 1: Set Up Your Project

First, create a new directory for your custom Roach build. Inside that directory, create a `main.go` file.

```sh
mkdir my-roach-custom
cd my-roach-custom
touch main.go
```

This `main.go` will be the entry point for your custom interpreter.

## Step 2: Write the Main Go File

Open `main.go` and add the following code. This will be our template. It imports the necessary Roach packages and sets up the structure for adding your own code.

```go
// my-roach-custom/main.go
package main

import (
	// You must import the Roach evaluator and object packages.
	"github.com/LabMarket/roach/evaluator"
	"github.com/LabMarket/roach/object"

	// We also import the Roach CLI to run the interpreter.
	// Note: This assumes the Roach project is structured to export its CLI.
	"github.com/LabMarket/roach/cli"

	// --- Add your desired Go packages here ---
	"fmt"
	"time"
)

// This init() function is where you will register your custom functions.
func init() {
	// The first argument is the name of the function as it will be called in Roach.
	// The second is the Go function that implements it.
	evaluator.RegisterBuiltin("time.now", timeNow)
}

// --- Define your wrapper functions here ---

// This is a "wrapper" function. It bridges the gap between the Roach interpreter
// and your Go code.
//
// It must always have the signature:
// func(env *object.Environment, args ...object.Object) object.Object
func timeNow(env *object.Environment, args ...object.Object) object.Object {

	// 1. Validate arguments (if any)
	if len(args) != 0 {
		return evaluator.newError("wrong number of arguments. got=%d, want=0", len(args))
	}

	// 2. Call the Go code
	now := time.Now().Format(time.RFC3339)

	// 3. Convert the Go result back to a Roach object and return it
	return &object.String{Value: now}
}


// The main function simply starts the Roach CLI.
func main() {
	cli.Main()
}
```

*A Note on Project Structure:* For this to work, the Roach project must be structured as a Go module that other projects can import. The `main` function of the standard Roach build would need to be extracted into an exportable function, which we've called `cli.Main()` here.

## Step 3: Build Your Custom Interpreter

Now, from inside your `my-roach-custom` directory, you need to initialize a Go module and then build your executable.

```sh
# 1. Initialize a Go module for your custom build
go mod init my-roach-custom

# 2. Tidy dependencies. This will download the Roach interpreter and any
#    other packages you imported.
go mod tidy

# 3. Build your custom executable
go build -o my-roach .
```

You will now have a new executable file named `my-roach` in your directory.

## Step 4: Use Your Custom Function

Your new `my-roach` binary works exactly like the standard one, but it now includes the `time.now` function you added.

Create a script to test it, e.g., `test.roach`:

```roach
#!/usr/bin/env roach

puts("Hello from a custom build!
");
let now = time.now();
puts("The current time is: ", now, "
");
```

Run it with your custom binary:

```sh
./my-roach test.roach
```

**Expected Output:**

```
Hello from a custom build!
The current time is: 2023-10-27T10:00:00Z
```

That's it! You have successfully extended Roach with a new function written in Go. You can follow this same pattern to bind functions from any third-party package like `github.com/anacrolix/torrent` or to add your own custom utilities.
