package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const(
	player = "omxplayer"
)

func main(){
	var musicDir string
	var enabled bool
	var timer int
	var playlistLen int

	flag.StringVar(&musicDir,"dir", "/home/pi/music", "music directory")
	flag.IntVar(&playlistLen,"length", 5, "music playlist length")
	flag.BoolVar(&enabled,"play", true, "player enabled")
	flag.IntVar(&timer,"playTime", 0, "play full playlist with minute timer")
	flag.Parse()

	if !enabled{
		if err := killPlayer(); err != nil{
			log.Printf("[WARN] %s\n", err)
		}
		return
	}

	isRunning := checkPlayerStatus()
	if isRunning{
		log.Printf("[WARN] no,no,no, one at a time. player is already running, exiting\n")
		return
	}

	useTimePlay := false
	if timer != 0 {
		useTimePlay = true
		log.Printf("[INFO] play stop time set for %d minutes\n", timer)
		log.Printf("[INFO] --ignoring playlist length of %d\n", playlistLen)
		go starKillTimer(time.Minute * time.Duration(timer))
	}

	playlist, err := createPlaylist(musicDir, playlistLen, useTimePlay)
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
