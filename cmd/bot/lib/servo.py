#!/usr/bin/env python3

import time
import RPi.GPIO as gpio

# middlepos = (leftpos - rightpos) / 2 + leftpos
default_config = {
    'servo_pin': 37,
    'frequency': 50,
    'calibration': {
        'left': 0.5,
        'right': 2.5
    }
}


class Servo:

    def __init__(self, config=default_config):
        if gpio.getmode() != gpio.BOARD:
            gpio.setmode(gpio.BOARD)
        self.msPerCylce = 1000 / config['frequency']
        gpio.setup(config['servo_pin'], gpio.OUT)
        self.pwm = gpio.PWM(config['servo_pin'], config['frequency'])

    def pos(self, deg):
        pos = ((leftpos - rightpos) / 180)
        return pos * deg + leftpos

    def move(self, deg):
        interval = self.pos((deg + 90) * -1)
        dutyPerc = interval * 100 / self.msPerCylce
        self.pwm.start(dutyPerc)
        time.sleep(0.5)
        self.pwm.stop()


if __name__ == "__main__":
    servo = Servo()
    positions = [-90, -45, 0, 45, 90]
    for p in positions:
        servo.move(p)
        time.sleep(0.5)
    gpio.cleanup()
