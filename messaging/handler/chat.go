package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/Alfeenn/messaging/constant"
	"github.com/Alfeenn/messaging/model"
	"github.com/novalagung/gubrak/v2"
)

var connections = model.NewSocketConnections().ConnectionSocket()

func HandleIO(currentConn *model.WebSocketConnection, connections []*model.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, constant.MESSAGE_NEW_USER, "")

	for {
		payload := model.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, constant.MESSAGE_LEAVE, "")
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, constant.MESSAGE_CHAT, payload.Message)
	}
}

func ejectConnection(currentConn *model.WebSocketConnection) {
	filtered := gubrak.From(connections).Reject(func(each *model.WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	connections = filtered.([]*model.WebSocketConnection)
}

func broadcastMessage(currentConn *model.WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(model.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
	}
}
