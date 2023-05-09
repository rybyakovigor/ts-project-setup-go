package editorconfig

import (
	"linters-configuration/pkg/core"
)

var config = `root = true
[*]
charset = utf-8
indent_style = space
indent_size = 2
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
max_line_length = off
trim_trailing_whitespace = false`

func createConfig() {
	core.CreateConfig(".editorconfig", config)
}

func Install() {
	createConfig()
}
