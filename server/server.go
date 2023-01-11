package server

import (
	"errors"
	"github.com/KikosS41/GoMPG/server/entities"
)

func AddNewPlayer(players entities.Players, newPlayer entities.Player) (entities.Players, error) {
	for _, player := range players.All {
		if player.Name == newPlayer.Name {
			return players, errors.New("this player already exist")
		}
	}

	players.All = append(players.All, newPlayer)
	return players, nil
}

func UpdatePlayer(players entities.Players, playerToUpdate entities.Player) (entities.Players, error) {
	for i, player := range players.All {
		if player.Name == playerToUpdate.Name {
			players.All[i].Coordinate["x"] = playerToUpdate.Coordinate["x"]
			players.All[i].Coordinate["y"] = playerToUpdate.Coordinate["y"]
			players.All[i].Coordinate["z"] = playerToUpdate.Coordinate["z"]
			players.All[i].AttackDamage = playerToUpdate.AttackDamage
			players.All[i].Health = playerToUpdate.Health
			return players, nil
		}
	}
	return players, errors.New("player is not connected")
}

func DeletePlayer(players entities.Players, playerToDelete entities.Player) (entities.Players, error) {
	for i, other := range players.All {
		if other.Name == playerToDelete.Name {
			players.All = append(players.All[:i], players.All[i+1:]...)
			return players, nil
		}
	}
	return players, errors.New("can't disconnect player to delete, not found")
}

func Server() {
	var players entities.Players
	responser(players)
}
