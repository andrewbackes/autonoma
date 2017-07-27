#!/usr/bin/env python3

from bot.bot import Bot
from bot.sensors.sensor import Sensor
from bot.motors.driver import Driver
from bot.motors.servo import Servo

from bot.emulation.courses import box
from bot.emulation.emulator import Emulator


def new_emulator():
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
    return Emu


def run_emulator():
    Emu = new_emulator()
    Emu.start()

if __name__ == "__main__":
    print("Running Emulator")
    run_emulator()
