#!/usr/bin/env python3

from stepper import Stepper
# from servo import Servo


class RoofMount:

    def __init__(self):
        self._stepper = Stepper()
        self._stepper.disable()
        # self._servo = Servo()

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

    def home(self):
        '''Move stepper and servo to home positions'''
        self._stepper.home()
        pass


if __name__ == "__main__":
    print("Roof mount self test.")
    roofmount = RoofMount()
    print("Full clockwise rotation....")
    for degrees in range(36):
        roofmount.clockwise(10)
    print("Done.")
    print("Full counter-clockwise rotation....")
    for degrees in range(36):
        roofmount.counter_clockwise(10)
    print("Done.")
