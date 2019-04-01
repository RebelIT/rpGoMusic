# rpGoMusic
RaspberryPi GOLang music player

## Whats this for?
My kids always want to listen to music when playing in the basement but fiddle too much with radio buttons and are not old enough to use an mp3
player with more than one button.  This will be a MP3 player with a GPIO button to play music.  Push a button and it will play music from a
local library of MP3's for 5 songs at a time. (can be override)

`GOOS=linux GOARCH=arm go build -o player .`
`scp player pi@10.0.0.xxx:/home/pi`


## Usage
* play 5 songs and quit:  `./player`
* play 10 songs and quit:  `./player -length 10`
* play 5 songs from a specific directory: `./player -dir "/media/music/kids"`
* play 8 songs from a specific directory: `./player -dir "/media/music/kids" -length 8`
* help:  `./player --help`


## More info
* hard coded to omxplayer for raspbian
* tested on Pi3b but should work on any others as long as supported audio is enabled


## ToDo
* document GPIO pin control & installation
* integration into [rpIoT](https://github.com/RebelIT/rpIoT)
* overrides via API to play based on playtime time vs. playlist length