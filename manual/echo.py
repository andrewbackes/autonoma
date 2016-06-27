#!/usr/bin/env python

import time
import RPi.GPIO as gpio

sensors = {
    'front' : {
        'echo': 16,
        'trigger': 12 },
    'back' : {
        'echo': 18,
        'trigger': 22 }
}

def distance(sensor, measure='cm'):
    gpio.setmode(gpio.BOARD)
    gpio.setup(sensor['trigger'], gpio.OUT)
    gpio.setup(sensor['echo'], gpio.IN)

    gpio.output(sensor['trigger'], True)
    time.sleep(0.00001) # specified wait time for this hardware
    gpio.output(sensor['trigger'], False)
    print "Ultrasonic sensor:"
    
    print gpio.input(sensor['echo'])

    while gpio.input(sensor['echo']) == 0:
        pass
        
    start = time.time()
    
    lapsed = time.time() - start
    while gpio.input(sensor['echo']) == 1:
        lapsed = time.time() - start
        if lapsed > 0.004:
            print "nothing in range"
            return None
    
    if measure == 'cm':
        distance = lapsed / 0.000058
    elif measure == 'in':
        distance = lapsed / 0.000148
    else:
        print 'unsupported measurement'
        distance = None
    gpio.cleanup()

    print str(distance) + ' ' + measure
    return distance


if __name__ == "__main__":
    try:
        for sensor in sensors:
            distance(sensor, 'in')
    except KeyboardInterrupt:
        gpio.cleanup()
