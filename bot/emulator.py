#!/usr/bin/env python3

from bot import Bot
from sensors.sensor import Sensor
from motors.driver import Driver
from motors.servo import Servo


def run_emulator():
    emulated_sensors = {
        "ultrasonic": Sensor,
        "irdistance": Sensor,
        "irproximity": Sensor,
        "compass": Sensor,
    }
    emulated_motors = {
        "driver": Driver,
        "servo": Servo,
    }
    Emulator = Bot(emulated_sensors, emulated_motors)
    Emulator.start()

if __name__ == "__main__":
    print("Running Emulator")
    run_emulator()
