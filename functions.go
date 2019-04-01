package main

import (
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func playSong(songPath string) error {
	if songPath == ""{
		//sometimes the random function inserts a blank line, quick short term fix for it.
		return nil
	}
	args := []string{}
	args = append(args, songPath)

	cmd := exec.Command(player, args...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[INFO] Waiting for song to finish\n")
	err = cmd.Wait()
	if err != nil{
		log.Printf("[ERROR] Song finished with error: %v\n", err)
		return err
	}
	return nil
}

func createPlaylist(dir string, numberOfSongs int) (playlist []string, error error) {
	log.Printf("[INFO] creating playlist\n")

	playlist, err := getAllSongs(dir)
	if err != nil{
		log.Printf("[ERROR] creating playlist\n")
		return nil, err
	}
	log.Printf("[INFO] playlist length: %v\n", len(playlist))
	randomize(playlist)
	return trim(playlist, numberOfSongs), nil
}

func getAllSongs(dir string) ([]string, error){
	log.Printf("[INFO] searching all songs in %s\n", dir)
	songs := []string{""}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("[ERROR] Unable to read music directory: %s\n",dir)
		log.Println(err)
		return nil, err
	}

	for _, f := range files {
		//log.Printf("[DEBUG] appending: %s\n", f.Name())
		songs = append(songs, f.Name())
	}

	//log.Printf("[DEBUG] raw song list length: %v\n", len(songs))
	return songs, nil
}

func randomize(songs []string) {
	log.Printf("[INFO] randomizing")
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for len(songs) > 0 {
		n := len(songs)
		randIndex := r.Intn(n)
		songs[n-1], songs[randIndex] = songs[randIndex], songs[n-1]
		songs = songs[:n-1]
	}
}

func trim(list []string, length int) []string {
	log.Printf("[INFO] trimming")
	newList := []string{}
	for i, p := range list{
		if i <= length -1 {
			log.Printf("[INFO] Adding song: %d -- %s\n", i, p)
			newList = append(newList, p)
		}
	}
	return newList
}

func checkPlayerStatus()(running bool){
	ps, _ := process.Processes()

	for _, p := range ps{
		name, _ := p.Name()
		if name == "omxplayer"{
			//log.Printf("[DEBUG] process name matches: %s\n", name)
			return true
		} else{
			//log.Printf("[DENUG] process name does not match: %s\n", name)
		}
	}

	return false
}

/* future overrides from flags
func killTimer(minutes time.Duration){
	time.Sleep(time.Minute * minutes)
}

func killPlayer() (output string, err error) {
	bytes, err := exec.Command("pkill", "omxplayer").Output()
	output = string(bytes)
	return
}
*/