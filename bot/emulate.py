#!/usr/bin/env python3

from bot import Bot
from sensors.sensor import Sensor
from motors.driver import Driver
from motors.servo import Servo

from emulation.courses import box
from emulation.emulator import Emulator


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
    Emu = Bot(emulated_sensors, emulated_motors)
    Emulator.occupied = box.occupied
    Emu.start()

if __name__ == "__main__":
    print("Running Emulator")
    run_emulator()
