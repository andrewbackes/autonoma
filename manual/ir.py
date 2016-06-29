#!/usr/bin/env python3

import RPi.GPIO as gpio

sensors = {
    'left': 32,
    'right': 36
}

def blocked(sensor):
    gpio.setmode(gpio.BOARD)
    gpio.setup(sensor, gpio.IN)
    return gpio.input(sensor) == 0
