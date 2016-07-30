#!/usr/bin/env python3

import RPi.GPIO as gpio
import time

sensors = {
    'left': 36,
    'right': 32
}

def blocked(sensor="all"):
    gpio.setmode(gpio.BOARD)
    if sensor == 'all':
        gpio.setup(sensors['left'], gpio.IN)
        gpio.setup(sensors['right'], gpio.IN)
        blocked = gpio.input(sensors['left']) == 0 and gpio.input(sensors['right']) == 0
    else:
        gpio.setup(sensors[sensor], gpio.IN)
        blocked = gpio.input(sensors[sensor]) == 0
    gpio.cleanup()
    return blocked
