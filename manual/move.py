#!/usr/bin/env python

import time
import RPi.GPIO as gpio
import readchar

def init():
    gpio.setmode(gpio.BOARD)
    gpio.setup(7, gpio.OUT)
    gpio.setup(11, gpio.OUT)
    gpio.setup(13, gpio.OUT)
    gpio.setup(15, gpio.OUT)

def toggle(a, b, c, d, t):
    init()
    gpio.output(7, a)
    gpio.output(11, b)
    gpio.output(13, c)
    gpio.output(15, d)
    time.sleep(t)
    gpio.cleanup()

def forward(t):
    toggle(False, True, True, False, t)

def backward(t):
    toggle(True, False, False, True, t)

def turn_left(t):
    toggle(False, True, False, False, t)

def turn_right(t):
    toggle(True, True, True, False, t)

def counter_clockwise(t):
    toggle(False, True, False, True, t)

def clockwise(t):
    toggle(True, False, True, False, t)


if __name__ == "__main__":
    counter_clockwise(0.1)
    clockwise(0.1)
