#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

#
# Nema 17
#
#  - Coil #1: Red & Yellow wire pair. Coil #2 Green & Brown/Gray wire pair.
#  - Datasheet:
#       https://cdn-shop.adafruit.com/product-files/324/C140-A+datasheet.jpg
#


class Stepper:

    CLOCKWISE = 1
    COUNTER_CLOCKWISE = 0

    _config = {
        "gpio": {
            "step":     36,
            "dir":      32,
            "ms1":      16,  # GPIO-23
            "ms2":      22,  # GPIO-25
            "ms3":      18,  # GPIO-24
            "enable":   12   # GPIO-18
        },
        "stepsPerRevolution": 200,  # 1.8 per step
        "stepDelay": 1.0 / 200,
        "resolution": {
            #        MS1,          MS2,           MS3
            'Full': (gpio.LOW,     gpio.LOW,      gpio.LOW),
            'Half': (gpio.HIGH,    gpio.LOW,      gpio.LOW),
            '1/4':  (gpio.LOW,     gpio.HIGH,     gpio.LOW),
            '1/8':  (gpio.HIGH,    gpio.HIGH,     gpio.LOW),
            '1/16': (gpio.HIGH,    gpio.HIGH,     gpio.HIGH),
        },
    }

    def __init__(self, config={}):
        self._config.update(config)
        gpio.setmode(gpio.BOARD)
        gpio.setup(self._config['gpio']['dir'], gpio.OUT)
        gpio.setup(self._config['gpio']['step'], gpio.OUT)
        gpio.output(self._config['gpio']['dir'], self.CLOCKWISE)
        gpio.setup(self._config['gpio']['enable'], gpio.OUT)
        # for micro-stepping:
        self._mode = (
            self._config['gpio']['m1'],
            self._config['gpio']['m2'],
            self._config['gpio']['m3']
        )
        gpio.setup(self._mode, gpio.OUT)
        gpio.output(self._mode, self._config['resolution']['1/16'])

    def enable(self):
        gpio.output(self._config['gpio']['enable'], gpio.LOW)

    def disable(self):
        gpio.output(self._config['gpio']['enable'], gpio.HIGH)

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
    gpio.cleanup()
