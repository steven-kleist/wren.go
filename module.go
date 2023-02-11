package main

import wren "github.com/crazyinfin8/WrenGo"

type builtinModule struct {
	source string
	module *wren.Module
	registered bool
}

func loadModuleCallback(vm *wren.VM, name string) (string, bool) {
	// check if name is an internal module
	if mod, ok := Modules[name]; ok {
		if !mod.registered {
			vm.SetModule(name, mod.module)
			mod.registered = true
		}
		return mod.source, true
	}

	// module [name] not found
	return "", false
}
