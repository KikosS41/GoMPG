package server

import (
	"errors"
	"github.com/KikosS41/GoMPG/server/entities"
)

func AddNewPlayer(players *[]entities.Player, newPlayer entities.Player) error {
	for _, player := range *players {
		if player.Name == newPlayer.Name {
			return errors.New("this player already exist")
		}
	}

	*players = append(*players, newPlayer)
	return nil
}

func DeletePlayer(players []entities.Player, playerToDelete entities.Player) ([]entities.Player, error) {
	for i, other := range players {
		if other.Name == playerToDelete.Name {
			return append(players[:i], players[i+1:]...), nil
		}
	}
	return players, errors.New("can't disconnect player to delete, not found")
}

func Server() {
	var players []entities.Player
	responser(players)
}
