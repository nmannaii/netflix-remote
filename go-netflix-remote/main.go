package main

import (
	"embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed icons/win&mac/icon.ico
var icon []byte
var app *App

func main() {
	// Create an instance of the app structure
	app = NewApp()
	go setupSystray()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-netflix-remote",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         app.startup,
		HideWindowOnClose: true,
		DisableResize:     true,
		OnDomReady:        app.domReady,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}

func setupSystray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Netflix Remote")
	systray.SetTooltip("Netflix Remote")

	expandItem := systray.AddMenuItem("Restore", "Restore windows")
	quitItem := systray.AddMenuItem("Quit", "Quit the application")
	go func() {
		<-quitItem.ClickedCh
		systray.Quit()
		runtime.Quit(app.ctx)
	}()
	go func() {
		<-expandItem.ClickedCh
		runtime.WindowShow(app.ctx)
	}()
}

func onExit() {
	// Clean up resources, if any
	runtime.Quit(app.ctx)
}
