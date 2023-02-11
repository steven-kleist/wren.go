package main

import (
	_ "embed"
	wren "github.com/crazyinfin8/WrenGo"
)

//go:embed module_io.wren
var module_io_source string

func init() {
	Modules["io"] = builtinModule{
		source: module_io_source,
		module: wren.NewModule(wren.ClassMap{}),
	}
}
