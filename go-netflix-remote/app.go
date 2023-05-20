package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	socketio "github.com/googollee/go-socket.io"
	"github.com/nmannaii/go-netflix-remote/controllers"
	"github.com/nmannaii/go-netflix-remote/utils"
	"github.com/skip2/go-qrcode"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Msg struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Init Gin
	r := gin.Default()
	server := socketio.NewServer(nil)
	controllers.GroupKeyPressRoutes(r)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		runtime.LogTrace(a.ctx, "connected:")
		return nil
	})

	server.OnEvent("/", "move-mouse", func(s socketio.Conn, msg Msg) {
		robotgo.DragSmooth(int(msg.X), int(msg.Y))
	})
	go func() {
		if err := server.Serve(); err != nil {
			runtime.LogError(ctx, fmt.Sprintf("socketio listen error: %s\n", err))
		}
		runtime.LogTrace(a.ctx, "Socket io listening")
	}()
	defer server.Close()
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	r.Static("/public", "./statics/netflix-remote")
	err := r.Run(":3698")
	if err != nil {
		runtime.LogError(ctx, "Error starting server.")
		return
	}
	runtime.LogTrace(a.ctx, "Server start on port 3698")
}

func (a *App) GetIpAddressQrCode() string {
	ip := utils.GetLocalIpAddress()
	png, _ := qrcode.Encode("http://"+ip+":3698", qrcode.Medium, 256)

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
}
