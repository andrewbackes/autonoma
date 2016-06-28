#!/usr/bin/env python3

import RPi.GPIO as gpio

def init():
    gpio.setmode(gpio.BOARD)
    gpio.setup(32, gpio.IN)

def blocked():
    return gpio.input(32) == 0
