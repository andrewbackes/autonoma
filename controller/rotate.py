#!/usr/bin/env python3

import time
import move
import heading

def diff(a, b):
    if a > b:
        return a - b
    return b - a

def direction(target, current):
    if (pos - degrees) % 360 > 180:
        return move.clockwise
    else:
        return move.counter_clockwise

def to_heading(degrees, accuracy=6):
    power = 80
    pos = heading.degrees()
    direction(degrees, pos)


    else:
        

    


'''

degrees = 10
pos = 350

(degrees - pos) % 360 = 20
(pos - degress) % 360 = 340



10 + 359 = 369 % 360 = 9

1 - 10 = 351

a + b = c

'''