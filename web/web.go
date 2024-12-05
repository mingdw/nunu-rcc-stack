package web

import "embed"

//go:embed index.html favicon.ico html asserts
var HtmlsFs embed.FS
