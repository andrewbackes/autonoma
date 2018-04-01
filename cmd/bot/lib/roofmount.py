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
        '''Move up relative to current position'''
        try:
            self._servo.set_position(self._servo.position - degrees)
        except:
            pass

    def down(self, degrees=10):
        '''Move down relative to current position'''
        try:
            self._servo.set_position(self._servo.position + degrees)
        except:
            pass

    def level(self):
        '''Vertically level the mount'''
        self._servo.set_position(self._config['servo']['level_degrees'])

    def clockwise(self, degrees=10):
        '''Move clockwise relative to current position'''
        self.__rotate(Stepper.CLOCKWISE, degrees)

    def counter_clockwise(self, degrees=10):
        '''Move counter-clockwise relative to current position'''
        self.__rotate(Stepper.COUNTER_CLOCKWISE, -degrees)

    def __rotate(self, dir, degrees):
        self._stepper.enable()
        self._stepper.set_direction(dir)
        pos = (self._stepper.position() + degrees) % 360
        self._stepper.set_position(pos)
        self._stepper.disable()

    def home(self):
        '''Move stepper and servo to home positions'''
        self._stepper.home()
        self.level()

    def horizontal_position(self):
        return self._stepper.position()

    def vertical_position(self):
        return self._servo.position()

    def position(self):
        return (self.horizontal_position(), self.vertical_position())

    def set_position(self, horizontal, vertical):
        # self._stepper.set_position(horizontal)
        self._servo.set_position(vertical)


def self_test():
    print("Roof mount self test.")
    roofmount = RoofMount()
    # Servo:
    print("Vertical movement test")
    roofmount.set_position(roofmount.horizontal_position(),
                           roofmount._config['servo']['min_degrees'])
    for degrees in range(11):
        roofmount.up()
    roofmount.level()
    print("Done.")
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
