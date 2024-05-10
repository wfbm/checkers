package server

import (
	"checkers/game"
	"fmt"
	"net"
)

var Port int

type Server struct {
	Address  string
	Rooms    []game.GameRoom
	listener net.Listener
}

func NewServer() Server {
	return Server{
		Address: fmt.Sprintf(":%d", Port),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	s.listener = ln
	go s.accept()

	return nil
}

func (s Server) accept() error {

	for {

		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnections(conn)
	}
}

func (s Server) handleConnections(conn net.Conn) {

}
