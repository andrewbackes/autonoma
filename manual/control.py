#!/usr/bin/env python

import RPi.GPIO as gpio
import readchar
from move import *

def repl():
    print "Use w,a,s,d to control the vehicle. x to exit"
    t = 0.2
    while True:
        k = readchar.readkey()
        if k == "w":
            forward(t)
        elif k == "s":
            backward(t)
        elif k == "a":
            pivot_left(t)
        elif k == "d":
            pivot_right(t)
        elif k == "x":
            break
    print "done"