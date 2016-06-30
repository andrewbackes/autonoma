#!/usr/bin/env python3

import time
import RPi.GPIO as gpio


GPIO.setmode(GPIO.BOARD)

# motors
gpio.setup(7, gpio.OUT)
gpio.setup(11, gpio.OUT)
gpio.setup(13, gpio.OUT)
gpio.setup(15, gpio.OUT)
gpio.output(7, False)
gpio.output(11, True)
gpio.output(13, True)
gpio.output(15, False)

# pwd
GPIO.setup(40, GPIO.OUT)
time.sleep(1.0)


p = GPIO.PWM(40, 50)
p.start(50)
time.sleep(2.0)

p.ChangeDutyCycle(90)
p.ChangeFrequency(100)

time.sleep(2.0)
p.stop()
GPIO.cleanup()