#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

servo_pin = 37

freq = 50

rightpos = 0.5
leftpos = 2.5
middlepos = (leftpos - rightpos) / 2 + leftpos

msPerCylce = 1000 / freq

pwm = None


def init():
    gpio.setup(servo_pin, gpio.OUT)
    pwm = gpio.PWM(servo_pin, freq)


def pos(deg):
    pos = ((leftpos - rightpos) / 180)
    return pos * deg + leftpos


def move(deg):
    interval = pos((deg + 90) * -1)
    dutyPerc = interval * 100 / msPerCylce
    if not pwm:
        print("must initialize servo.init()")
        return
    pwm.start(dutyPerc)
    time.sleep(0.5)
    pwm.stop()


if __name__ == "__main__":
    gpio.setmode(gpio.BOARD)
    init()
    positions = [-90, -45, 0, 45, 90]
    for p in positions:
        move(p)
        time.sleep(0.5)
    gpio.cleanup()
