#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
import move
import servo
#import heading
#import ir
from getch import *

def repl():
    print("Use w,a,s,d to control the vehicle. x to exit")
    t = 0.2
    speed = 50
    step = 10
    servo_pos = 0
    servo.move(servo_pos)
    while True:
        #print(str(heading.degrees()) + "°" + " @ " + str(speed) + "% power.")
        print(str(speed) + "% power.")
        #print("Sensors blocked - left: " + str(ir.blocked('left')) + " right: " + str(ir.blocked('right')) )#+ " SONIC: " + str(d) + "")
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
        elif k == 'f':
            speed = max(0, speed - step)
        elif k == "q":
            servo_pos = max(-90, servo_pos - 45)
            servo.move(servo_pos)
        elif k == "e":
            servo_pos = min(90, servo_pos + 45)
            servo.move(servo_pos)
        elif k == "x":
            break
        print("servo pos: " + str(servo_pos) + "°")
    print("done")

if __name__ == "__main__":
    repl()