package game

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type Player struct {
	ID         string
	Connection net.Conn
}

type GameRoom struct {
	ID      int64
	Players []Player
}

func NewRoom() GameRoom {

	roomId := time.Now().Unix()
	room := GameRoom{
		ID: roomId,
	}

	return room
}

func (r GameRoom) IsAvailable() bool {
	return len(r.Players) < 2
}

func (r *GameRoom) EnterRoom(player Player) error {

	if !r.IsAvailable() {
		return errors.New("not possible for connectiong into this room, it is currently not available")
	}

	r.Players = append(r.Players, player)

	return nil
}

func (r *GameRoom) LeaveRoom(playerToLeave Player) error {

	newRoomPlayerList := []Player{}

	for _, player := range r.Players {
		if player.ID != playerToLeave.ID {
			newRoomPlayerList = append(newRoomPlayerList, player)
		}
	}

	r.Players = newRoomPlayerList

	return nil
}

func NewPlayer(conn net.Conn) Player {
	player := Player{}
	player.ID = fmt.Sprintf("p-%d", time.Now().Unix())
	player.Connection = conn
	return player
}

func (p Player) SendCommand(cmd string) {

}
