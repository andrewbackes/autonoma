#!/usr/bin/evn python3

from bot.sensors.sensor import Sensor
from bot.motors.driver import Driver
from bot.motors.servo import Servo

import math

import logging

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)


def frange(start, stop, step):
    i = start
    while i < stop:
        yield i
        i += step


class Emulator(Sensor):

    # Static variables used to share state between Driver and Sensors
    location = (0, 0)
    heading = 0.0
    occupied = set({})

    def __init__(self, metadata, config):
        super().__init__(metadata, config)
        logger.debug("Initializing Emulator.")
        logger.debug(self.metadata)
        Emulator.location = (0, 0)
        Emulator.heading = 0.0

    def _read_compass(self):
        return Emulator.heading

    def _read_distance(self):
        logger.debug("reading distance")
        start_angle = (Emulator.heading +
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
                logger.debug("Checking angle: %d" % a)
                rad = ((a % 360) * math.pi) / 180
                x = Emulator.location[0] + math.floor(d * math.cos(rad))
                y = Emulator.location[1] + math.floor(d * math.sin(rad))
                logger.debug("Checking point (%d, %d)" % (x, y))
                if (x, y) in Emulator.occupied:
                    dist = math.sqrt(
                        (x - Emulator.location[0])**2 +
                        (y - Emulator.location[1])**2
                    )
                    return dist
        return 0.0

    def read(self):
        if 'compass' in self.metadata['id']:
            return self._read_compass()
        return self._read_distance()

    def face(self, heading):
        Emulator.heading = heading
        return

    def move(self, distance, location):
        Emulator.location = location
        return
