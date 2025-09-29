package cli

import (
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/LabMarket/roach/evaluator"
	"github.com/LabMarket/roach/lexer"
	"github.com/LabMarket/roach/object"
	"github.com/LabMarket/roach/parser"
)

// IsCustom should be set to true by custom builds in an init() function.
// This controls whether the CLI calls os.Exit or returns an error.
var IsCustom = false

// This version-string will be updated via CI system for generated binaries.
var version = "master/unreleased"

//go:embed data/stdlib.roach
var stdlib string

//
// Implemention of "version()" function.
//
func versionFun(args ...object.Object) object.Object {
	return &object.String{Value: version}
}

//
// Implemention of "args()" function.
//
func argsFun(args ...object.Object) object.Object {
	l := len(os.Args[1:])
	result := make([]object.Object, l)
	for i, txt := range os.Args[1:] {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

//
// Execute the supplied string as a program.
//
func Execute(input string) error {

	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		var sb strings.Builder
		for _, msg := range p.Errors() {
			sb.WriteString(fmt.Sprintf("\t%s\n", msg))
		}
		return fmt.Errorf(sb.String())
	}

	// Register a function called version()
	// that the script can call.
	evaluator.RegisterBuiltin("version",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (versionFun(args...))
		})

	// Access to the command-line arguments
	evaluator.RegisterBuiltin("args",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (argsFun(args...))
		})

	//
	//  Parse and evaluate our standard-library.
	//
	initL := lexer.New(stdlib)
	initP := parser.New(initL)
	initProg := initP.ParseProgram()
	evaluator.Eval(initProg, env)

	//
	//  Now evaluate the code the user wanted to load.
	//
	//  Note that here our environment will still contain
	// the code we just loaded from our data-resource
	//
	//  (i.e. Our roach-based standard library.)
	//
	evaluator.Eval(program, env)
	return nil
}

func Main() error {

	//
	// Setup some flags.
	//
	eval := flag.String("eval", "", "Code to execute.")
	vers := flag.Bool("version", false, "Show our version and exit.")

	//
	// Parse the flags
	//
	flag.Parse()

	//
	// Showing the version?
	//
	if *vers {
		fmt.Printf("roach %s\n", version)
		return nil
	}

	//
	// Executing code?
	//
	if *eval != "" {
		return Execute(*eval)
	}

	//
	// Otherwise we're either reading from STDIN, or the
	// named file containing source-code.
	//
	var input []byte
	var err error

	if len(flag.Args()) > 0 {
		input, err = ioutil.ReadFile(os.Args[1])
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		return fmt.Errorf("Error reading: %s", err.Error())
	}

	return Execute(string(input))
}
