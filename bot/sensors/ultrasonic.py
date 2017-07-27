#!/usr/bin/env python3

from sensors.sensor import Sensor


class Ultrasonic(Sensor):

    def read(self):
        return 0.0
