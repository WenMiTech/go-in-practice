package main

import (
	"fmt"
	"music_player/libs"
	"bufio"
	"os"
	"strings"
	"music_player/mp"
	"strconv"
)

var lib *libs.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i);
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&libs.Music{strconv.Itoa(id),
				tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}

	case "remove":
		if len(tokens) == 3 {
			index, err := strconv.Atoi(tokens[2]);
			if err != nil {
				fmt.Println(err)
			}
			lib.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])

	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {
	//var musicManager *libs.MusicManager = libs.NewMusicManager()
	//var music *libs.Music = &libs.Music{"1", "child", "Taylor", "child.map", "popular"}

	//musicManager := libs.NewMusicManager()
	//music := &libs.Music{"1", "child", "Taylor", "child.map", "popular"}
	//musicManager.Add(music);
	//fmt.Println(musicManager)
	fmt.Println(`
        Enter following commands to control the player:
        lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)
	lib = libs.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command -> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0]);
		}
	}

}
