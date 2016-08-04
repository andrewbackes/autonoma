#!/usr/bin/env python3

import time
import RPi.GPIO as gpio


servo_pin = 37

gpio.setmode(gpio.BOARD)

gpio.setup(servo_pin, gpio.OUT)
freq = 50
pwm = gpio.PWM(servo_pin, freq)

leftpos = 0.75
rightpos = 2.5
middlepos = (rightpos - leftpos) / 2 + leftpos

poslist = [leftpos, middlepos, rightpos, middlepos]

msPerCylce = 1000 / freq


for i in range(3):
    for pos in poslist:
        dutyPerc = pos * 100 / msPerCylce
        print("Pos: " + str(pos))
        print("Duty: " + str(dutyPerc) )
        print()
        pwm.start(dutyPerc)
        time.sleep(0.5)

pwm.stop()
gpio.cleanup()