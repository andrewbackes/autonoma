#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

servo_pin = 37

freq = 50

rightpos = 0.5
leftpos = 2.5
middlepos = (leftpos - rightpos) / 2 + leftpos

msPerCylce = 1000 / freq

def pos(deg):
    pos = ((leftpos - rightpos) / 180)
    return pos*deg + leftpos

def move(deg):
    interval = pos((deg + 90)*-1)
    gpio.setmode(gpio.BOARD)
    gpio.setup(servo_pin, gpio.OUT)
    pwm = gpio.PWM(servo_pin, freq)
    dutyPerc = interval * 100 / msPerCylce
    pwm.start(dutyPerc)
    time.sleep(0.5)
    pwm.stop()
    gpio.cleanup()
