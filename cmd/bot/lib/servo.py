#!/usr/bin/env python3

import time
import pigpio


class Servo:
    _config = {
        'gpio': 37,
        'gpioBCN': 26,
        'frequency': 50,
        'ratio': 1,
        'secondsPer60deg': 0.19,
        'calibration': {
            'right': 0.5,
            'left': 2.5
        }
    }
    __pos = 0
    __pi = None

    def __init__(self, config=None):
        if config:
            self._config.update(config)
        print("Servo config: ", config)
        self.msPerCylce = 1000 / self._config['frequency']

        self.__pi = pigpio.pi()
        if not self.__pi.connected:
            print("Could not connect to pigpiod.")
        self.__pi.set_mode(self._config['gpioBCN'], pigpio.OUTPUT)
        self.move(0)  # move to center position

    def __del__(self):
        self.__pi.stop()

    def __calc_interval(self, deg):
        pos = ((self._config['calibration']['left'] -
                self._config['calibration']['right']) / 180)
        return pos * deg + self._config['calibration']['left']

    def move(self, deg):
        # adjust for a possible external gear ratio:
        adjusted_deg = deg / self._config['ratio']
        interval = self.__calc_interval((adjusted_deg + 90) * -1)
        dutyPerc = interval * 100 / self.msPerCylce

        # self.__pi.set_PWM_frequency(
        #    self._config['gpioBCN'], self._config['frequency'])
        # self.__pi.set_PWM_dutycycle(self._config['gpioBCN'], dutyPerc)
        # time.sleep(0.5 + self.__spin_time(deg))
        # self.__pi.set_PWM_dutycycle(self._config['gpioBCN'], 0)

        self.__pi.set_servo_pulsewidth(17, 0)
        time.sleep(1)
        self.__pi.set_servo_pulsewidth(17, 1000)
        time.sleep(1)
        self.__pi.set_servo_pulsewidth(17, 1500)
        time.sleep(1)
        self.__pi.set_servo_pulsewidth(17, 2000)
        time.sleep(1)
        self.__pi.set_servo_pulsewidth(17, 0)
        self.__pos = deg

    def position(self):
        return self.__pos

    def __spin_time(self, deg):
        if deg > self.__pos:
            diff = deg - self.__pos
        else:
            diff = self.__pos - deg
        return (self._config['secondsPer60deg'] * (diff / 60))


if __name__ == "__main__":
    servo = Servo()

#    pi.set_servo_pulsewidth(17, 0)    # off
#    pi.set_servo_pulsewidth(17, 1000) # safe anti-clockwise
#    pi.set_servo_pulsewidth(17, 1500) # centre
#    pi.set_servo_pulsewidth(17, 2000) # safe clockwise
    positions = [-90, -45, 0, 45, 90, 0]
    for p in positions:
        print("Position ", p)
        servo.move(p)
        time.sleep(0.5)
