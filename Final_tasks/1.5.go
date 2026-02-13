package main

import (
	"slices"
	"fmt"
)

type Player struct {
	Name string
	Goals int
	Misses int
	Assists int
	Rating float64
}

func calculateRating(goals, misses, assists int) float64 {
	var rating float64
	if misses == 0{
		rating = float64(goals) + float64(assists) / 2
	} else {
		rating = (float64(goals) + (float64(assists) / 2)) / float64(misses)
	}
	return rating
}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: calculateRating(goals, misses, assists)}
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Goals > b.Goals || (a.Goals == b.Goals && a.Name > b.Name):
			return 1
		case a.Goals < b.Goals || (a.Goals == b.Goals && a.Name < b.Name):
			return -1
		default:
			return 0
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Rating > b.Rating || (a.Rating == b.Rating && a.Name > b.Name):
			return 1
		case a.Rating < b.Rating || (a.Rating == b.Rating && a.Name < b.Name):
			return -1
		default:
			return 0
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case float64(a.Goals) / float64(a.Misses) > float64(b.Goals) / float64(b.Misses) || (float64(a.Goals) / float64(a.Misses) == float64(b.Goals) / float64(b.Misses) && a.Name > b.Name):
			return 1
		case float64(a.Goals) / float64(a.Misses) < float64(b.Goals) / float64(b.Misses) || (float64(a.Goals) / float64(a.Misses) == float64(b.Goals) / float64(b.Misses) && a.Name < b.Name):
			return -1
		default:
			return 0
		}
	})
	return players
}

func main() {
	players := []Player{
		NewPlayer("a", 10, 5, 3),
		NewPlayer("b", 15, 7, 2),
		NewPlayer("c", 8, 2, 5),
	}

	fmt.Println(goalsSort(players))
}