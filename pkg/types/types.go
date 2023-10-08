package types

import (
	"linters-configuration/pkg/core"
)

func Install() {
	install()
}

func install() {
	core.WriteScript("check:types", "tsc --noEmit")
}
