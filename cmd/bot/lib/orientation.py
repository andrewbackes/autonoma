#!/usr/bin/env python3

from Adafruit_BNO055 import BNO055
import time


class Orientation:

    def __init__(self):
        # use i2c
        self.bno055 = BNO055.BNO055()

    def heading(self):
        time.sleep(0.01)
        heading, roll, pitch = self.bno055.read_euler()
        return heading


if __name__ == "__main__":
    orientation = Orientation()
    while True:
        print(orientation.heading())
        time.sleep(1)

    # Other values you can optionally read:
    # Orientation as a quaternion:
    # x,y,z,w = bno.read_quaterion()
    # Sensor temperature in degrees Celsius:
    # temp_c = bno.read_temp()
    # Magnetometer data (in micro-Teslas):
    # x,y,z = bno.read_magnetometer()
    # Gyroscope data (in degrees per second):
    # x,y,z = bno.read_gyroscope()
    # Accelerometer data (in meters per second squared):
    # x,y,z = bno.read_accelerometer()
    # Linear acceleration data (i.e. acceleration from movement, not gravity--
    # returned in meters per second squared):
    # x,y,z = bno.read_linear_acceleration()
    # Gravity acceleration data (i.e. acceleration just from gravity--returned
    # in meters per second squared):
    # x,y,z = bno.read_gravity()
    # Sleep for a second until the next reading.
