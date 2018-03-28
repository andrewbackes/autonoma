#!/usr/bin/env python3

import time
import RPi.GPIO as gpio


class Servo:
    _config = {
        'gpio': 37,
        'frequency': 50,
        'ratio': 1,
        'secondsPer60deg': 0.19,
        'calibration': {
            'right': 0.5,
            'left': 2.5
        }
    }
    __pos = 0

    def __init__(self, config=None):
        if config:
            self._config.update(config)
        print("Servo config: ", config)
        if gpio.getmode() != gpio.BOARD:
            gpio.setmode(gpio.BOARD)
        self.msPerCylce = 1000 / self._config['frequency']
        gpio.setup(self._config['gpio'], gpio.OUT)
        self.move(0)  # move to center position

    def __calc_interval(self, deg):
        pos = ((self._config['calibration']['left'] -
                self._config['calibration']['right']) / 180)
        return pos * deg + self._config['calibration']['left']

    def move(self, deg):
        # adjust for a possible external gear ratio:
        adjusted_deg = deg / self._config['ratio']

        interval = self.__calc_interval((adjusted_deg + 90) * -1)
        dutyPerc = interval * 100 / self.msPerCylce
        pwm = gpio.PWM(self._config['gpio'], self._config['frequency'])
        pwm.start(dutyPerc)

        time.sleep(0.1 + self.__spin_time(deg))
        pwm.stop()
        self.__pos = deg

    def position(self):
        return self.__pos

    def __spin_time(self, deg):
        if deg > self.__pos:
            diff = deg - self.__pos
        else:
            diff = self.__pos - deg
        return (self._config['secondsPer60deg'] * (self.__delta(deg) / 60))


if __name__ == "__main__":
    servo = Servo()
    positions = [-90, -45, 0, 45, 90, 0]
    for p in positions:
        print("Position ", p)
        servo.move(p)
        time.sleep(0.5)
    gpio.cleanup()
