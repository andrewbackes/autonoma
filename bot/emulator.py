#!/usr/bin/env python3

from bot import Bot
from sensors.sensor import Sensor
from motors.driver import Driver
from motors.servo import Servo

from emulation.courses import box
from emulation.sensor import SensorEmulator


class Emulator(Bot):

    def __init__(self, occupied_map):
        emulated_sensors = {
            "ultrasonic": SensorEmulator,
            "irdistance": SensorEmulator,
            "irproximity": SensorEmulator,
            "compass": SensorEmulator,
        }
        emulated_motors = {
            "driver": SensorEmulator,
            "servo": SensorEmulator,
        }
        super().__init__(emulated_sensors, emulated_motors)
        SensorEmulator.occupied = occupied_map

if __name__ == "__main__":
    print("Running Emulator")
    emu = Emulator(box.occupied)
    emu.start()
