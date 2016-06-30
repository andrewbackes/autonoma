#!/usr/bin/env python3

import time
import RPi.GPIO as gpio


gpio.setmode(gpio.BOARD)

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
gpio.setup(40, gpio.OUT)
gpio.setup(38, gpio.OUT)

time.sleep(1.0)

pl = gpio.PWM(38, 50)
pr = gpio.PWM(40, 50)

pl.start(50)
pr.start(50)

time.sleep(1.0)

pl.ChangeDutyCycle(90)
pl.ChangeFrequency(100)
pr.ChangeDutyCycle(90)
pr.ChangeFrequency(100)

time.sleep(1.0)

pr.stop()
pl.stop()

gpio.cleanup()