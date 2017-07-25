#!/usr/bin/env python3

#
# For reference check out:
#   https://github.com/adafruit/Adafruit_Python_LSM303/blob/master/Adafruit_LSM303/LSM303.py
#

from sensors.sensor import Sensor


class Compass(Sensor):

    def read(self):
        return 0.0
