package main

import (
	"flag"
	"hangman-classic/CreateList"
	"hangman-classic/Game"
	"hangman-classic/RequestUsr"
	"hangman-classic/StartAndStop"
)

func init() {
	const (
		Save  = "Save.txt"
		usage = "init du jeu"
	)
	flag.String("SaveType", Save, usage)
}

func main() {
	if len(CreateList.ReadFile("save.txt")) >= 1 {
		SavedGame := RequestUsr.Level(true)
		for SavedGame == "" {
			SavedGame = RequestUsr.Level(true)
		}
		Game.Game(StartAndStop.Start())
	} else {
		var ListLetterUsed []string
		IndexOfDeath, LineHangman := 0, 0
		lev := RequestUsr.Level(false)
		for lev == "" {
			lev = RequestUsr.Level(false)
		}
		ListWord := CreateList.ReadFile(lev)
		ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
		DashList := CreateList.CreateDashList(ListWordCap)
		Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
	}
}
