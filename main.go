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

type Params struct{
	MusicDir	string
	Enabled		bool
	Timer 		int
	PlaylistLen	int
	Statsd 		string
}

var config *Params

func main(){
	config = setConfig()

	statStartProgram()
	//check force stop flag set to false
	if !config.Enabled{
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
	if config.Timer != 0 {  //use a play timer vs. static playlist length
		useTimePlay = true
		log.Printf("[INFO] play stop time set for %d minutes\n", config.Timer)
		log.Printf("[INFO] --ignoring playlist length of %d\n", config.PlaylistLen)
		go starKillTimer(config.Timer)
	}

	//do the work to create aplaylist
	playlist, err := createPlaylist(config.MusicDir, config.PlaylistLen, useTimePlay)
	if err != nil{
		fmt.Println(err)
		return
	}

	//play it
	timeStart := time.Now()
	for i, p := range playlist{

		log.Printf("[INFO] playing %d of %d: %s\n", i+1, len(playlist),p)
		statSongPlay(p)
		if err := playSong(config.MusicDir +"/"+ p); err != nil{
			break
		}
	}
	timeEnd := time.Now()
	statRuntime(timeDiff(timeStart, timeEnd))
	log.Printf("[INFO] playlist complete :yay:\n\n")
}

func setConfig()(config *Params){
	c := &Params{}

	//Parse input flags and sets defaults
	flag.StringVar(&c.MusicDir,"dir", "/home/pi/music", "music directory")
	flag.IntVar(&c.PlaylistLen,"length", 5, "music playlist length")
	flag.BoolVar(&c.Enabled,"play", true, "player enabled")
	flag.IntVar(&c.Timer,"playTime", 0, "play full playlist with minute timer")
	flag.StringVar(&c.Statsd,"statsdHost", "", "StatsdHost to emit mettics to")
	flag.Parse()

	return c
}
