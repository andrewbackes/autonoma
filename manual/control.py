#!/usr/bin/env python

import RPi.GPIO as gpio
import readchar
from move import *
import math

def repl():
    print "Use w,a,s,d to control the vehicle. x to exit"
    t = 0.2
    speed = 50
    step = 10
    while True:
        k = readchar.readkey()
        if k == "w":
            forward(t, speed)
        elif k == "s":
            backward(t, speed)
        elif k == "a":
            counter_clockwise(t, speed)
        elif k == "d":
            clockwise(t, speed)
        elif k == 'r':
            speed = math.max(0, speed - step)
        elif k == 'f':
            speed = math.min(100, speed + step)
        elif k == "x":
            break
    print "done"

if __name__ == "__main__":
    repl()