package main

import (
	"server/app/server"
)

func main() {
	s := server.Server{}
	if err := s.Start(); err != nil {
		panic(err)
	}
}
