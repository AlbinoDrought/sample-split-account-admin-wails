package main

import (
	"embed"
	"os"
	"runtime"

	"github.com/hectane/go-acl"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	fixedTokenKey = "SAMPLE_RANDOM_KEY"
	fixedTokenVal = "with-fixed-token"
	webviewDir    = "C:\\ProgramData\\Sample"
)

func main() {
	if runtime.GOOS == "windows" && os.Getenv(fixedTokenKey) != fixedTokenVal {
		runWithFixedToken()
	}

	println("Setting data dir to", webviewDir)
	if err := os.MkdirAll(webviewDir, os.ModePerm); err != nil {
		println("Failed creating dir:", err)
	}
	if err := acl.Chmod(webviewDir, 0777); err != nil {
		println("Failed setting ACL on dir:", err)
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sample-data-dir",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewUserDataPath: webviewDir,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
