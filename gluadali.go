package gluadali

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

type daliModule struct {
}

func NewDaliModule() *daliModule {
	return &daliModule{}
}

func (h *daliModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"sendGear": h.sendGear,
	})
	L.Push(mod)
	return 1
}

func (h *daliModule) sendgear(L *lua.LState) int {
	L.Push(lua.LNumber(L.ToInt(1)))
	fmt.Printf("sending opcode '%d' to device at addr %d\n", L.ToInt(1), L.ToInt(2))
	return 1
}
