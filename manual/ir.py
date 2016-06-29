#!/usr/bin/env python3

import RPi.GPIO as gpio
import time

sensors = {
    'left': 32,
    'right': 36
}

def blocked(sensor):
    gpio.setmode(gpio.BOARD)
    gpio.setup(sensor, gpio.IN)
    return gpio.input(sensor) == 0


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