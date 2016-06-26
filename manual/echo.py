#!/usr/bin/env python

import time
import RPi.GPIO as gpio

#pin number
echo = 16       # gpio 23
trigger = 12    # gpio 15

def distance(measure='cm'):
    gpio.setmode(gpio.BOARD)
    gpio.setup(trigger, gpio.OUT)
    gpio.setup(echo, gpio.IN)

    gpio.output(trigger, True)
    time.sleep(0.00001) # specified wait time for this hardware
    gpio.output(trigger, False)
    print "triggering sensor"
    
    print gpio.input(echo)

    while gpio.input(echo) == 0:
        pass
        
    start = time.time()
    
    while gpio.input(echo) == 1 and lapsed = time.time() - start < 0.004:
        pass
    
    if measure == 'cm':
        distance = lapsed / 0.000058
    elif measure == 'in':
        distance = lapsed / 0.000148
    else:
        print 'unsupported measurement'
        distance = None
    gpio.cleanup()

    return distance


if __name__ == "__main__":
    try:
        print distance('in')
    except KeyboardInterrupt:
        gpio.cleanup()
