package main

import (
	"embed"

	"github.com/circa10a/witchonstephendrive.com/cmd"
)

//go:embed web
var frontendAssets embed.FS

//go:embed api
var apiDocAssets embed.FS

// @title witchonstephendrive.com
// @version 0.1.0
// @description Control my halloween decorations
// @contact.name Caleb Lemoine
// @contact.email caleblemoine@gmail.com
// @license.name MIT
// @license.url https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/main/LICENSE
// @host witchonstephendrive.com
// @BasePath /
// @Schemes https
func main() {
	cmd.Execute(frontendAssets, apiDocAssets)
}
