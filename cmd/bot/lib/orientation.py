#!/usr/bin/env python3

from Adafruit_BNO055 import BNO055

# use i2c
bno055 = BNO055.BNO055()


def heading():
    heading, roll, pitch = bno055.read_euler()
    return heading


if __name__ == "__main__":
    while True:
        # Read the Euler angles for heading, roll, pitch (all in degrees).
        heading, roll, pitch = bno055.read_euler()
        # Read the calibration status, 0=uncalibrated and 3=fully calibrated.
        sys, gyro, accel, mag = bno.get_calibration_status()
        print('Heading={0:0.2F} Roll={1:0.2F} Pitch={2:0.2F}\tSys_cal={3} ' +
              'Gyro_cal={4} Accel_cal={5} Mag_cal={6}'.format(
                  heading, roll, pitch, sys, gyro, accel, mag))

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
