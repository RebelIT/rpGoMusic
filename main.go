package main

import (
	"flag"
	"fmt"
	"log"
)

const(
	player = "omxplayer"
)

func main(){
	var musicDir string
	var playlistLen int
	flag.StringVar(&musicDir,"dir", "/home/pi/music", "music directory")
	flag.IntVar(&playlistLen,"length", 5, "music playlist length")
	flag.Parse()

	isRunning := checkPlayerStatus()
	if isRunning{
		log.Printf("[WARN] no,no,no, one at a time. player is already running, exiting\n")
		return
	}

	playlist, err := createPlaylist(musicDir, playlistLen)
	if err != nil{
		fmt.Println(err)
		return
	}

	for i, p := range playlist{
		log.Printf("[INFO] playing %d of %d: %s\n", i, playlistLen,p)
		if err := playSong(musicDir +"/"+ p); err != nil{
			break
		}
	}
	log.Printf("[INFO] playlist complete :yay:\n\n")
}
