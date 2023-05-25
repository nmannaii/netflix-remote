package main

import (
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/nmannaii/fyne-netflix-remote/backend"
	"github.com/nmannaii/fyne-netflix-remote/utils"
	"github.com/skip2/go-qrcode"
	"image/color"
	"strconv"
)

func main() {

	a := app.NewWithID("io.netflix-remote")
	a.SetIcon(resourceIconIco)
	w := a.NewWindow("Hello")
	w.SetTitle("Netflix Remote")
	w.SetIcon(resourceIconIco)
	mainContainer := container.NewVBox()
	trayMenu := fyne.NewMenu(
		"Menu",
		fyne.NewMenuItem("Show", func() {
			w.Show()
		}),
	)
	a.(desktop.App).SetSystemTrayIcon(resourceIconIco)
	a.(desktop.App).SetSystemTrayMenu(trayMenu)
	setRichTxt(mainContainer)
	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(400, 300))
	w.SetFixedSize(true)

	server := backend.Server()
	go server.InitGin()
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}

func setRichTxt(ctr *fyne.Container) {
	ipAddress := utils.GetLocalIpAddress()
	img := canvas.NewImageFromResource(resourceNetflixRemotePng)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(400, 120))
	rich := widget.NewRichTextFromMarkdown(`
# Welcome to Netflix Remote

An opensource free , simple and user-friendly app, that you can use it on IOS or Android.

You need only three things:
* Your favorite web browser (preferred Chrome)
* Netflix navigator plugin: https://www.netflixnavigator.com/ to navigate using arrows
* Being connected to the same network
* Now open your browser and access this url: **http://` + ipAddress + `:` + strconv.Itoa(utils.PORT) + `**
`)
	qrCode := canvas.NewImageFromReader(bytes.NewReader(getQrCode()), "qrCode")
	qrCode.FillMode = canvas.ImageFillContain
	qrCode.SetMinSize(fyne.NewSize(250, 250))
	scanMe := canvas.NewText("Scan Me", color.Black)
	scanMe.TextSize = 16
	scanMe.Alignment = fyne.TextAlignCenter
	ctr.Add(container.NewCenter(img))
	ctr.Add(rich)
	ctr.Add(container.NewCenter(qrCode))
	ctr.Add(scanMe)
}

func getQrCode() []byte {
	ip := utils.GetLocalIpAddress()
	png, _ := qrcode.Encode("http://"+ip+":3698", qrcode.Medium, 256)

	return png
}
