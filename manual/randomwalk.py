#!/usr/bin/env python3

import time
import RPi.GPIO as gpio
import move
import echo

increment = 0.3
unit = 'in'
range = 15
delay = 0.3

def spin_detect(positions):
    distances = [None] * positions
    for pos in range(positions/2):
        d = echo.distances(unit)
        distances[pos] = d['front']
        distances[pos * 2] = d['back']
        move.clockwise(increment)
        time.sleep(0.03)
    return distances


def move_until_blocked(dist):
    print("moving")
    while True:
        d = echo.distance(echo.sensors['front'], unit)
        print("Nearest object: " + str(d)) 
        if d and d <= dist:
            print("BLOCKED")
            return
        move.forward(increment)
        time.sleep(delay)


def find_unblocked_path(dist, clockwise, counterclockwise):
    print("rotating")
    open_positions = 0
    while open_positions < 5:
        clockwise(increment)
        d = echo.distance(echo.sensors['front'], unit)
        if not d or d < dist*2:
            open_positions += 1
        else:
            open_positions = 0
        time.sleep(delay)
    print("FOUND PATH")
    counterclockwise(increment * 3)
    

def walk():
    counter = 0
    funcs = [move.clockwise, move.counter_clockwise]
    while True:
        move_until_blocked(range)
        time.sleep(0.1)
        find_unblocked_path(range, funcs[counter % 2], funcs[ (counter+1) % 2])
        time.sleep(0.1)
        counter += 1


if __name__ == "__main__":
    walk() 