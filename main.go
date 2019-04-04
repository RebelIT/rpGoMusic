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
	var enabled bool
	var timer int
	var playlistLen int

	//Parse input flags and sets defaults
	flag.StringVar(&musicDir,"dir", "/home/pi/music", "music directory")
	flag.IntVar(&playlistLen,"length", 5, "music playlist length")
	flag.BoolVar(&enabled,"play", true, "player enabled")
	flag.IntVar(&timer,"playTime", 0, "play full playlist with minute timer")
	flag.Parse()

	//check force stop flag set to false
	if !enabled{
		if err := killPlayer(); err != nil{
			log.Printf("[WARN] %s\n", err)
		}
		return
	}

	//check if omxplayer is already playing
	isRunning := checkPlayerStatus()
	if isRunning{
		log.Printf("[WARN] no,no,no, one at a time. player is already running, exiting\n")
		return
	}


	useTimePlay := false
	if timer != 0 {  //use a play timer vs. static playlist length
		useTimePlay = true
		log.Printf("[INFO] play stop time set for %d minutes\n", timer)
		log.Printf("[INFO] --ignoring playlist length of %d\n", playlistLen)
		go starKillTimer(timer)
	}

	//do the work to create aplaylist
	playlist, err := createPlaylist(musicDir, playlistLen, useTimePlay)
	if err != nil{
		fmt.Println(err)
		return
	}

	//play it
	for i, p := range playlist{
		log.Printf("[INFO] playing %d of %d: %s\n", i+1, len(playlist),p)
		if err := playSong(musicDir +"/"+ p); err != nil{
			break
		}
	}
	log.Printf("[INFO] playlist complete :yay:\n\n")
}
