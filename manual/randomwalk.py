#!/usr/bin/env python3

import time
import RPi.GPIO as gpio
import move
import echo
import ir

spin_increment = 0.2
forward_increment = 0.1
unit = 'in'
shortest_dist = 12
delay = 0.01

"""
def spin_detect(positions):
    distances = [None] * positions
    for pos in range(int(positions/2)):
        d = echo.distances(unit)
        print(pos)
        print(d['front'])
        distances[pos] = d['front']

        b = int(positions/2) + pos
        print(b)
        print(d['back'])
        distances[ b ] = d['back']

        move.clockwise(increment)
        time.sleep(0.03)
    return distances
"""


def move_until_blocked(dist):
    print("moving")
    while True:
        d = echo.distance(echo.sensors['front'], unit)
        blocked = d and d <= dist
        if ir.blocked() or blocked: 
            print("BLOCKED - (IR: " + str(ir.blocked()) + " SONIC: " + str(blocked) + ")")
            return
        move.forward(forward_increment)
        time.sleep(delay)
    move.backward(forward_increment * 2)


"""
def find_unblocked_path(dist, clockwise, counterclockwise):
    print("rotating")
    open_positions = 0
    while open_positions < 5:
        clockwise(increment)
        d = echo.distance(echo.sensors['front'], unit)
        far_enough = not d or d < dist*2
        if not ir.blocked() and far_enough:
            open_positions += 1
        else:
            open_positions = 0
        time.sleep(delay)
    print("FOUND PATH")
    counterclockwise(increment * 3)
"""    

def find_unblocked_path(dist, clockwise, counterclockwise):
    print("rotating")
    start = time.time()
    while ir.blocked() or echo.blocked(shortest_dist):
        clockwise(spin_increment)
        #time.sleep(0.001)
    lapsed = time.time() - start
    clockwise(lapsed)


def walk():
    counter = 0
    funcs = [move.clockwise, move.counter_clockwise]
    while True:
        move_until_blocked(shortest_dist)
        time.sleep(0.1)
        find_unblocked_path(shortest_dist, funcs[counter % 2], funcs[ (counter+1) % 2])
        time.sleep(0.1)
        counter += 1
    gpio.cleanup()


if __name__ == "__main__":
    try:
        walk() 
    except KeyboardInterrupt:
        gpio.cleanup()