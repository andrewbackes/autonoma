#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

#
# Nema 17
#  Note:
#  - Coil #1: Red & Yellow wire pair. Coil #2 Green & Brown/Gray wire pair.
#


class Stepper:

    CLOCKWISE = 1
    COUNTER_CLOCKWISE = 0

    _config = {
        "gpio": {
            "step": 36,
            "dir": 32
        },
        "stepsPerRevolution": 200,  # 1.8 per step
        "stepDelay": 1.0 / 200,
    }

    def __init__(self, config={}):
        self._config.update(config)
        gpio.setmode(gpio.BOARD)
        gpio.setup(self._config['gpio']['dir'], gpio.OUT)
        gpio.setup(self._config['gpio']['step'], gpio.OUT)
        gpio.output(self._config['gpio']['dir'], self.CLOCKWISE)

    def one(self):
        for x in range(self._config['stepsPerRevolution']):
            gpio.output(self._config['gpio']['step'], gpio.HIGH)
            time.sleep(self._config['stepDelay'])
            gpio.output(self._config['gpio']['step'], gpio.LOW)
            time.sleep(self._config['stepDelay'])


if __name__ == "__main__":
    print("Stepper test")
    stepper = Stepper()
    stepper.one()
