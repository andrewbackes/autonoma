#!/usr/bin/env python3

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

    while gpio.input(sensor['echo']) == 0:
        pass
        
    start = time.time()
    
    lapsed = time.time() - start
    while gpio.input(sensor['echo']) == 1:
        lapsed = time.time() - start
        if lapsed > 0.004:
            return None
    
    if measure == 'cm':
        distance = lapsed / 0.000058
    elif measure == 'in':
        distance = lapsed / 0.000148
    else:
        raise ValueError('unsupported measurement')
        distance = None
    gpio.cleanup()

    return distance

def distances(measure="in"):
    distances = {}
    for sensor in sensors:
        distances[sensor] = distance(sensors[sensor], 'in')
    print(distances)
    return distances

def blocked(dist):
    d = distance(echo.sensors['front'], unit)
    bl = d and d <= dist
    return bl

if __name__ == "__main__":
    try:
        distances('in')
    except KeyboardInterrupt:
        gpio.cleanup()
