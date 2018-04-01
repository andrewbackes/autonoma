#!/usr/bin/env python3

from stepper import Stepper
from servo import Servo


class RoofMount:

    _config = {
        'servo': {
            'level_degrees': 35,
            'max_degrees': 70,   # down
            'min_degrees': -48,  # up
        }
    }

    def __init__(self, config={}):
        self._config.update(config)
        self._stepper = Stepper()
        self._stepper.disable()
        self._servo = Servo()

    def up(self, degrees=10):
        pass

    def down(self, degrees=10):
        pass

    def clockwise(self, degrees=10):
        self.__rotate(Stepper.CLOCKWISE, degrees)

    def counter_clockwise(self, degrees=10):
        self.__rotate(Stepper.COUNTER_CLOCKWISE, -degrees)

    def __rotate(self, dir, degrees):
        self._stepper.enable()
        self._stepper.set_direction(dir)
        pos = (self._stepper.position() + degrees) % 360
        self._stepper.set_position(pos)
        self._stepper.disable()

    def level(self):
        '''Vertically level the mount'''
        self._servo.set_position(self._config['servo']['level_degrees'])

    def home(self):
        '''Move stepper and servo to home positions'''
        self._stepper.home()
        pass


def self_test():
    print("Roof mount self test.")
    roofmount = RoofMount()
    # Servo:
    roofmount.level()
    return

    # Stepper:
    print("Full clockwise rotation....")
    for degrees in range(36):
        roofmount.clockwise(10)
    print("Done.")
    print("Full counter-clockwise rotation....")
    for degrees in range(36):
        roofmount.counter_clockwise(10)
    print("Done.")

if __name__ == "__main__":
    self_test()
