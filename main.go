package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	wren "github.com/crazyinfin8/WrenGo"
)

var (
	Version string                   = "0.1.0"
	Modules map[string]builtinModule = map[string]builtinModule{}
)

func ShowVersion() {
	fmt.Printf(
		"%s v.%s (%s)",
		filepath.Base(os.Args[0]),
		Version,
		runtime.Version(),
	)
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Shows the current version.")
	flag.Parse()

	if showVersion {
		ShowVersion()
		os.Exit(0)
	}

	var filename string
	if filename = flag.Arg(0); filename == "" {
		fmt.Printf("Please provide a filepath.")
		os.Exit(1)
	}

	cfg := wren.NewConfig()
	cfg.LoadModuleFn = loadModuleCallback
	vm := cfg.NewVM()
	defer vm.Free()

	if err := vm.InterpretFile(filename); err != nil {
		fmt.Printf("Error while executing script: %v\n", err.Error())
		os.Exit(1)
	}
}
