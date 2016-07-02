#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

pwm = {
    'left': 38,
    'right': 40
}

driver = {
    'in1': 15,
    'in2': 13,
    'in3': 11,
    'in4': 77,
}

def init():
    gpio.setmode(gpio.BOARD)
    gpio.setup(7, gpio.OUT)
    gpio.setup(11, gpio.OUT)
    gpio.setup(13, gpio.OUT)
    gpio.setup(15, gpio.OUT)
    gpio.setup(40, gpio.OUT)
    gpio.setup(38, gpio.OUT)
    gpio.setup(pwm['left'], gpio.OUT)
    gpio.setup(pwm['right'], gpio.OUT)

def run_at_power(t, power=80):
    p = {}
    p['left'] = gpio.PWM(pwm['left'], 100) # 100 hz
    p['right'] = gpio.PWM(pwm['right'], 100)
    p['left'].start(power)
    p['right'].start(power)
    time.sleep(t)
    p['left'].stop()
    p['right'].stop()

def toggle(a, b, c, d, t, power=80):
    init()
    gpio.output(7, a)
    gpio.output(11, b)
    gpio.output(13, c)
    gpio.output(15, d)
    run_at_power(t, power)
    gpio.cleanup()

def forward(t, power=80):
    toggle(False, True, True, False, t, power)

def backward(t, power=80):
    toggle(True, False, False, True, t, power)

def turn_left(t, power=80):
    toggle(False, True, False, False, t, power)

def turn_right(t, power=80):
    toggle(True, True, True, False, t, power)

def counter_clockwise(t, power=80):
    toggle(False, True, False, True, t, power)

def clockwise(t, power=80):
    toggle(True, False, True, False, t, power)


if __name__ == "__main__":
    counter_clockwise(0.1)
    clockwise(0.1)
