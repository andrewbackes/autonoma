#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

default_config = {
    'pwm': {
        'left': 38,
        'right': 40
    },
    'driver': {
        'in4': 7,
        'in3': 11,
        'in2': 13,
        'in1': 15,
    }
}


class Move:

    def __init__(self, config=default_config):
        self.config = config
        print("H-bridge config:", config)
        if gpio.getmode() != gpio.BOARD:
            gpio.setmode(gpio.BOARD)
        for pin in self.config['pwm'].values():
            gpio.setup(value, gpio.OUT)
        for pin in self.config['driver'].values():
            gpio.setup(value, gpio.OUT)

    def __run_at_power(self, t, power=80):
        p = {}
        p['left'] = gpio.PWM(self.config['pwm']['left'], 100)  # 100 hz
        p['right'] = gpio.PWM(self.config['pwm']['right'], 100)
        p['left'].start(power)
        p['right'].start(power)
        time.sleep(t)
        p['left'].stop()
        p['right'].stop()

    def __toggle(self, a, b, c, d, t, power=80):
        gpio.output(self.config['driver']['in4'], a)
        gpio.output(self.config['driver']['in3'], b)
        gpio.output(self.config['driver']['in2'], c)
        gpio.output(self.config['driver']['in1'], d)
        self.__run_at_power(t, power)

    def forward(self, t, power=80):
        print("forward @ " + str(power) + " for t=" + str(t))
        self.toggle(False, True, True, False, t, power)
        print("movement done")

    def backward(self, t, power=80):
        print("backward @ " + str(power) + " for t=" + str(t))
        self.toggle(True, False, False, True, t, power)
        print("movement done")

    def turn_left(self, t, power=80):
        self.toggle(False, True, False, False, t, power)

    def turn_right(self, t, power=80):
        self.toggle(True, True, True, False, t, power)

    def counter_clockwise(self, t, power=80):
        print("counter_clockwise @ " + str(power) + " for t=" + str(t))
        self.toggle(False, True, False, True, t, power)
        print("movement done")

    def clockwise(self, t, power=80):
        print("clockwise @ " + str(power) + " for t=" + str(t))
        self.toggle(True, False, True, False, t, power)
        print("movement done")


if __name__ == "__main__":
    move = Move()
    move.counter_clockwise(0.1)
    move.clockwise(0.1)
    gpio.cleanup()
