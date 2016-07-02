#!/usr/bin/env python

import RPi.GPIO as gpio
import readchar
from move import *
import heading

def repl():
    print "Use w,a,s,d to control the vehicle. x to exit"
    t = 0.2
    speed = 50
    step = 10
    while True:
        print heading.degrees() + "\370"
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
            speed = min(100, speed + step)
            print speed
        elif k == 'f':
            speed = max(0, speed - step)
            print speed
        elif k == "x":
            break
    print "done"

if __name__ == "__main__":
    repl()