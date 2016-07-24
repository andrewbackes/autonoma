#!/usr/bin/env python3
# -*- coding: utf-8 -*-

print("starting")

import time
import RPi.GPIO as gpio
import move
import echo
import ir
import heading
from random import randint

spin_increment = 0.2
spin_power = 80

forward_increment = 0.1
forward_power = 50
forward_delay = 0.05

unit = 'in'
shortest_dist = 12
delay = 0.2


def move_until_blocked(dist):
    print("moving forward @ " + str(heading.degrees()) + "°" )
    while True:
        l, r, sf = ir.blocked('left'), ir.blocked('right'), echo.blocked(shortest_dist, unit)
        print("Sensors blocked - left: " + str(l) + " right: " + str(r) + " sonic-front: " + str(sf))
        if l or r or sf: #or blocked: 
            print("blocked")
            return
        move.forward(forward_increment, forward_power)
        time.sleep(forward_delay)
    

def spin_direction():
    print("finding safe direction")
    l, r, sf = ir.blocked('left'), ir.blocked('right'), echo.blocked(shortest_dist, unit)
    if l and not r:
        return move.clockwise
    if r and not l:
        return move.counter_clockwise
    funcs = [move.clockwise, move.counter_clockwise]
    r = randint(0,1)
    print("choosing random spin direction (" + str(r) + ")")
    return funcs[r]
    
def face_unblocked_path(dist):
    print("rotating")
    start = time.time()
    while True:
        l, r, sf = ir.blocked('left'), ir.blocked('right'), echo.blocked(shortest_dist, unit)
        print("Sensors blocked - left: " + str(l) + " right: " + str(r) + " sonic-front: " + str(sf))
        spin = spin_direction()
        spin(spin_increment)
        time.sleep(0.1)
        if not l and not r and not sf:
            break
    lapsed = time.time() - start
    print("spin time: " + str(lapsed))
    spin = spin_direction()
    spin(lapsed)


def walk():
    funcs = [move.clockwise, move.counter_clockwise]
    while True:
        move_until_blocked(shortest_dist)
        time.sleep(delay)
        spin = spin_direction()
        time.sleep(delay)
        #move.backward(forward_increment * 2, forward_power)
        face_unblocked_path(shortest_dist)
        print("found unblocked path")
        time.sleep(delay)
    print("exit walk - cleaning up")
    gpio.cleanup()


if __name__ == "__main__":
    try:
        walk() 
    except KeyboardInterrupt:
        gpio.cleanup()
