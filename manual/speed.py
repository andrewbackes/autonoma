#!/usr/bin/env python3

import time
import move
import heading

def diff(a, b):
    if a > b:
        return a - b
    return b - a

def spin_speed(t, power):
    attempts = 3
    sum = 0
    for _ in range(attempts):
        time.sleep(0.1)
        h1 = heading.degrees()
        move.clockwise(t, power)
        h2 = heading.degrees()
        sum += diff(h1, h2)
    return sum/attempts

if __name__ == "__main__":
    print(spin_speed(0.1, 80))