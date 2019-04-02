import RPi.GPIO as GPIO
import time
import subprocess

GPIO.setmode(GPIO.BCM)

GPIO.setup(18, GPIO.IN, pull_up_down=GPIO.PUD_UP)

while True:
    input_state = GPIO.input(18)
    if input_state == False:
        print('Start Playing Music')
        subprocess.Popen(["/etc/apps/player"])
        time.sleep(0.2)