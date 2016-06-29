#!/usr/bin/env python3

import RPi.GPIO as gpio
import time

sensors = {
    'left': 36,
    'right': 32
}

def blocked():
    gpio.setmode(gpio.BOARD)
    gpio.setup(sensors['left'], gpio.IN)
    gpio.setup(sensors['right'], gpio.IN)
    return gpio.input(sensors['left']) == 0 and gpio.input(sensors['right']) == 0 


def dashboard():
    while True:
        l = blocked(sensors['left'])
        r = blocked(sensors['right'])
        block = {
            'left': l,
            'right': r
        }
        print(block)
        time.sleep(0.5)