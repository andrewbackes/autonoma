#!/usr/bin/env python

import time
import RPi.GPIO as gpio
import move
import echo

# hardcode these for now:
increment = 0.1
positions = 16
unit = 'in'

# learn this value on startup
speed = None

def spin_detect():
    distances = [None] * positions
    for pos in range(positions/2):
        d = echo.distances(unit)
        distances[pos] = d['front']
        distances[pos * 2] = d['back']
        move.clockwise(increment)
        time.sleep(0.03)
    return distances


def furthest_blocked_pos():
    distances = spin_detect()
    max = 0
    max_pos = None
    for pos in range(len(distances)): 
        dist = distances[pos]
        if dist and dist > max:
            max = dist
            max_pos = pos 
    # face forward
    move.clockwise(increment*(positions/2))
    return max_pos

def face(pos):
    # assume facing forward for now
    move.clockwise(increment*(pos))


def detect_speed():
    before = echo.distance(echo.sensors['front'])
    print before
    move.forward(increment)
    after = echo.distance(echo.sensors['front'])
    print after
    # hmmm.... not yet sure what to return on error
    try:
        return (before - after)/increment
    except:
        return None


def unblocked_position():
    for pos in range(positions):
        print "something"
    return None


def three_sixty():
    face(positions)

if __name__ == "__main__":
    print "Surroundings:"
    print spin_detect()
    
    #print "Detecting speed"
    #print "Patrolling..."
    #print "Speed: " + str(detect_speed())
    
    # loop:
        # spin in a circle
            # record front/back sensors
            # break when front sensor traces the back one
        # spin to the open path
            # move until until a wall is detected
