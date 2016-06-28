#!/usr/bin/env python3

import time
import RPi.GPIO as gpio
import move
import echo

increment = 0.1
unit = 'in'
range = 10

def move_until_blocked(dist):
    print("moving")
    while True:
        d = echo.distance(echo.sensors['front'], unit)
        print("Nearest object: " + str(d)) 
        if d and d <= dist:
            return
        move.forward(increment)


def find_unblocked_path(dist):
    print("rotating")
    open_positions = 0
    while open_positions < 3:
        move.clockwise(increment)
        d = echo.distance(echo.sensors['front'], unit)
        if not d or d < dist*2:
            open_positions += 1
        else:
            open_positions = 0
    move.counter_clockwise(increment)
    

def walk():
    while True:
        move_until_blocked(range)
        time.sleep(0.1)
        find_unblocked_path(range)
        time.sleep(0.1)


if __name__ == "__main__":
    walk() 