#!/usr/bin/evn python3

from sensors.sensor import Sensor
from motors.driver import Driver
from motors.servo import Servo

import math

import logging

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)


class SensorEmulator(Sensor):

    # Static variables used to share state between Driver and Sensors
    location = (0, 0)
    heading = 0.0
    occupied = set({})
    servo_position = 0

    def __init__(self, metadata, config):
        super().__init__(metadata, config)
        logger.debug("Initializing SensorEmulator.")
        logger.debug(self.metadata)

    def _read_compass(self):
        return SensorEmulator.heading

    def _read_distance(self):
        logger.debug("reading distance")
        start_angle = (SensorEmulator.heading +
                       self.metadata['angleOffset'] -
                       (self.metadata['coneWidth'] / 2)) % 360
        # can't mod 360 or the loop will end
        stop_angle = (start_angle + self.metadata['coneWidth'])

        logger.debug("Cone angle: %d to %d" % (start_angle, stop_angle))

        for d in frange(self.metadata['minDistance'],
                        self.metadata['maxDistance'],
                        1.0):
            logger.debug("Checking distance: %d" % d)

            for a in frange(start_angle, stop_angle, 0.25):
                logger.debug("Checking heading: %d" % a)
                polar = (-a + 90) % 360
                rad = (polar * math.pi) / 180
                logger.debug(
                    "Checking heading: %d   (Polor angle: %d degrees / %d radians)" %
                    (a, polar, rad))
                x = SensorEmulator.location[0] + math.floor(d * math.cos(rad))
                y = SensorEmulator.location[1] + math.floor(d * math.sin(rad))
                logger.debug("Checking point (%d, %d)" % (x, y))
                if (x, y) in SensorEmulator.occupied:
                    dist = math.sqrt(
                        (x - SensorEmulator.location[0])**2 +
                        (y - SensorEmulator.location[1])**2
                    )
                    return dist
        return 0.0

    def read(self):
        if 'compass' in self.metadata['id']:
            return self._read_compass()
        return self._read_distance()

    def rotate(self, heading):
        SensorEmulator.heading = heading
        return

    def position(self, degrees):
        SensorEmulator.servo_position = degrees
        return

    def move(self, distance, location):
        SensorEmulator.location = location
        return


def frange(start, stop, step):
    i = start
    while i < stop:
        yield i
        i += step
