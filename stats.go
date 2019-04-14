package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
)

func sendCounter(measurement string, tags statsd.Option){
	var c = config
	if c.Statsd == ""{
		log.Println("ignoring counter metric")
		return
	}
	addrOpt := statsd.Address(c.Statsd)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	s, err := statsd.New(addrOpt,fmtOpt,tags)
	if err != nil {
		log.Print(err)
	}
	defer s.Close()

	s.Increment(measurement)
}

func sendGauge(measurement string, tags statsd.Option, value int){
	var c = config
	if c.Statsd == ""{
		log.Println("ignoring histogram metric")
		//metrics disabled don't do anything
		return
	}
	addrOpt := statsd.Address(c.Statsd)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	s, err := statsd.New(addrOpt,fmtOpt,tags)
	if err != nil {
		log.Print(err)
	}
	defer s.Close()
	s.Gauge(measurement, value)
}

func statSongPlay(song string){
	//emits a new counter for every song played
	tags := statsd.Tags("song", song, "app", "player")
	measurement := "player_song"
	sendCounter(measurement,tags)
}

func statStartProgram(){
	//emits a new counter for every time the button is pushed
	tags := statsd.Tags("app", "player")
	measurement := "player_started"
	sendCounter(measurement,tags)
}

func statError(function string){
	//emits a new counter for every error in the program
	tags := statsd.Tags("function",function, "app", "player")
	measurement := "player_error"
	sendCounter(measurement,tags)
}

func statRuntime(seconds int){
	//emits a new histogram for total runtime
	tags := statsd.Tags("app", "player")
	measurement := "player_runtime"
	sendGauge(measurement,tags, seconds)
}