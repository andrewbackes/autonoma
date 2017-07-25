#!/usr/bin/env python3

from sensors.ultrasonic import Ultrasonic
from sensors.compass import Compass
from sensors.irdist import IRDistance
from sensors.irprox import IRProximity

from sensors.config import settings as sensor_settings

from motors.driver import Driver
from motors.servo import Servo

from motors.config import settings as motor_settings


class Bot(object):

    def __init__(self, sensor_factory, motor_factory):
        print("Initializing Bot.")
        self._register_sensors(sensor_factory)
        self._register_motors(motor_factory)
        self.set_location(0, 0)

    def _register_sensors(self, sensor_factory):
        default = {
            "ultrasonic": Ultrasonic,
            "irdistance": IRDistance,
            "irproximity": IRProximity,
            "compass": Compass,
        }
        self.sensors = {}
        self._register(
            sensor_factory,
            default,
            sensor_settings.items(),
            self.sensors)

    def _register_motors(self, motor_factory):
        default = {
            "driver": Driver,
            "servo": Servo,
        }
        self.motors = {}
        self._register(
            motor_factory,
            default,
            motor_settings.items(),
            self.motors)

    def _register(self, factory, default_factory, items, member_map):
        member_factory = factory if factory else default_factory
        for id, item in items:
            member_map[id] = member_factory[item['type']](
                item['metadata'], item['config'])
            print("Registered", item)

    def set_location(self, x, y):
        self.location_x = x
        self.location_y = y

    def start(self):
        print("Bot started.")

        print("Bot stopped.")
