#!/usr/bin/env python3

import time
import os
import pigpio


class Servo:
    _config = {
        'gpio': 37,
        'gpioBCN': 26,
        'frequency': 50,
        'ratio': 1,
        'secondsPer60deg': 0.19,
        'calibration': {
            'right': 500,
            'left': 2500
        }
    }
    __pos = 0
    __pi = None

    def __init__(self, config=None):
        if config:
            self._config.update(config)
        print("Servo config: ", config)

        self.__pi = pigpio.pi()
        if not self.__pi.connected:
            print("Could not connect to pigpiod.")
            os.exit(1)
        self.__pi.set_mode(self._config['gpioBCN'], pigpio.OUTPUT)
        self.set_position(0)  # move to center position

    def __del__(self):
        self.__pi.stop()

    def __calc_pulse_width(self, deg):
        pos = (self._config['calibration']['left'] -
               self._config['calibration']['right']) / 180
        return pos * (deg + 90) + self._config['calibration']['right']

    def __spin_time(self, deg):
        if deg > self.__pos:
            diff = deg - self.__pos
        else:
            diff = self.__pos - deg
        return (self._config['secondsPer60deg'] * (diff / 60))

    def set_position(self, deg):
        if deg > 75 or deg < -75:
            raise ValueError("Must be between -75 and 75")
        self.__pi.set_servo_pulsewidth(
            self._config['gpioBCN'], self.__calc_pulse_width(deg))
        time.sleep(0.1 + self.__spin_time(deg))
        self.__pi.set_servo_pulsewidth(self._config['gpioBCN'], 0)
        self.__pos = deg

    def position(self):
        return self.__pos


def self_test():
    servo = Servo()
    for p in range(-75, 75, 10):
        print("Position ", p)
        servo.set_position(p)
        time.sleep(0.5)
    del(servo)

if __name__ == "__main__":
    self_test()
