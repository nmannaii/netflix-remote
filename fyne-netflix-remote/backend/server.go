package backend

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	socketio "github.com/googollee/go-socket.io"
	"github.com/nmannaii/fyne-netflix-remote/backend/controllers"
	"github.com/nmannaii/fyne-netflix-remote/utils"
)

//go:embed netflix-remote/*
var remoteUi embed.FS

var LOGGER = log.Default()

var _server = &server{}

type server struct {
	SocketServer *socketio.Server
}

type coor struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func Server() *server {
	return _server
}

// InitGin init gin and websocket
func (s *server) InitGin() {
	r := gin.Default()
	s.SocketServer = socketio.NewServer(nil)
	controllers.GroupKeyPressRoutes(r)

	s.SocketServer.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("============= CONNECTED ===========")
		return nil
	})

	s.SocketServer.OnEvent("/", "move-mouse", func(s socketio.Conn, msg string) {
		obj := coor{}
		json.Unmarshal([]byte(msg), &obj)
		x, y := robotgo.GetMousePos()
		robotgo.Move(int(obj.X)+x, int(obj.Y)+y)
	})

	go func() {
		if err := s.SocketServer.Serve(); err != nil {
			LOGGER.Fatalln("socketio listen error: ", err)
		}
		LOGGER.Println("Socket io listening")
	}()

	defer s.SocketServer.Close()
	r.GET("/socket.io/*any", gin.WrapH(s.SocketServer))
	r.POST("/socket.io/*any", gin.WrapH(s.SocketServer))

	r.Use(static.Serve("/", utils.EmbedFolder(remoteUi, "netflix-remote", true)))
	err := r.Run(":3698")
	if err != nil {
		LOGGER.Fatalln("Error starting server.")
		return
	}
	LOGGER.Println("Server start on port 3698")
}
