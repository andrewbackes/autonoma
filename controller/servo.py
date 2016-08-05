#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

servo_pin = 37

freq = 50

leftpos = 0.25
rightpos = 2.5
middlepos = (rightpos - leftpos) / 2 + leftpos

poslist = [leftpos, middlepos, rightpos, middlepos]

positions = {
    "left" : leftpos,
    "middle" : middlepos,
    "right": rightpos }

msPerCylce = 1000 / freq

'''
def move(pos):
    gpio.setmode(gpio.BOARD)
    gpio.setup(servo_pin, gpio.OUT)
    pwm = gpio.PWM(servo_pin, freq)
    
    dutyPerc = positions[pos] * 100 / msPerCylce
    pwm.start(dutyPerc)
    time.sleep(0.5)
    pwm.stop()
    gpio.cleanup()
'''

def pos(deg):
    pos = (rightpos - leftpos) / 180
    return pos*deg + leftpos

def move(deg):
    interval = pos(deg + 90)
    gpio.setmode(gpio.BOARD)
    gpio.setup(servo_pin, gpio.OUT)
    pwm = gpio.PWM(servo_pin, freq)
    dutyPerc = interval * 100 / msPerCylce
    pwm.start(dutyPerc)
    time.sleep(0.5)
    pwm.stop()
    gpio.cleanup()

    

'''

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

'''
