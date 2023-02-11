package main

import (
	_ "embed"
	"os"
	"os/exec"
	"runtime"
	"strings"

	wren "github.com/crazyinfin8/WrenGo"
)

//go:embed module_os.wren
var module_os_source string

func init() {
	Modules["os"] = builtinModule{
		source: module_os_source,
		module: wren.NewModule(wren.ClassMap{
			"Platform": wren.NewClass(nil, nil, wren.MethodMap{
				"static name": os_platform_name,
			}),
			"Process": wren.NewClass(nil, nil, wren.MethodMap{
				"static allArguments": os_Process_allArguments,
				"static cwd":          os_Process_cwd,
				"static exec(_)":      os_Process_exec,
				"static exit_(_)":     os_Process_exit,
			}),
		}),
	}
}

func os_platform_name(
	vm *wren.VM,
	parameters []interface{},
) (interface{}, error) {
	return runtime.GOOS, nil
}

func os_Process_allArguments(
	vm *wren.VM,
	parameters []interface{},
) (interface{}, error) {
	list, _ := vm.NewList()
	for _, val := range os.Args {
		list.Insert(val)
	}

	return list, nil
}

func os_Process_cwd(
	vm *wren.VM,
	parameters []interface{},
) (interface{}, error) {
	if p, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return p, nil
	}
}

func os_Process_exec(
	vm *wren.VM,
	parameters []interface{},
) (interface{}, error) {
	cmdline := parameters[1].(string)
	params := strings.SplitN(cmdline, " ", 2)
	cmd := exec.Command(params[0], params[1])
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return cmd.ProcessState.ExitCode(), nil
}

func os_Process_exit(
	vm *wren.VM,
	parameters []interface{},
) (interface{}, error) {
	exit_code := parameters[1].(float64)
	os.Exit(int(exit_code))

	return nil, nil
}
