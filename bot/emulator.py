#!/usr/bin/env python3

import logging

from bot import Bot
from sensors.sensor import Sensor
from motors.driver import Driver
from motors.servo import Servo

from emulation.courses import box
from emulation.sensor import SensorEmulator

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)


class Emulator(Bot):

    def __init__(self, occupied_map):
        logger.info("Initializing Emulator.")
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
    logger.info("Running Emulator.")
    emu = Emulator(box.occupied)
    emu.start()
