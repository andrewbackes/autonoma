#!/usr/bin/env python3

import RPi.GPIO as gpio
import move
import heading
from getch import *

def repl():
    print "Use w,a,s,d to control the vehicle. x to exit"
    t = 0.2
    speed = 50
    step = 10
    while True:
        print(str(heading.degrees()) + "\370" + " @ " + str(speed) + "% power.")
        k = getch()
        if k == "w":
            move.forward(t, speed)
        elif k == "s":
            move.backward(t, speed)
        elif k == "a":
            move.counter_clockwise(t, speed)
        elif k == "d":
            move.clockwise(t, speed)
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