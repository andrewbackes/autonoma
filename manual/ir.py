#!/usr/bin/env python3

import RPi.GPIO as gpio

def blocked():
    gpio.setmode(gpio.BOARD)
    gpio.setup(32, gpio.IN)
    return gpio.input(32) == 0
