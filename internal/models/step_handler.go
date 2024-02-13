package models

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type StepHandler struct {
	ID   string
	Name string
	Used bool
}

func NewStepHandler(f interface{}) StepHandler {
	fp := runtime.FuncForPC(reflect.ValueOf(f).Pointer())

	return StepHandler{
		ID:   handlerKey(fp),
		Name: handlerName(fp),
		Used: false,
	}
}

func handlerKey(f *runtime.Func) string {
	file, line := f.FileLine(0)
	return fmt.Sprintf("%s:%s:%d", f.Name(), file, line)
}

func handlerName(f *runtime.Func) string {
	var pkg string
	name := f.Name()
	if idx := strings.LastIndex(name, "."); idx != -1 {
		pkg = name[:idx]
		name = name[idx+1:]
	}
	name = strings.Replace(name, "·", ".", -1)

	if pkg != "" {
		return fmt.Sprintf("%s:%s", pkg, name)
	}
	return name
}
