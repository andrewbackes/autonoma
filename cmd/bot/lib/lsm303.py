#!/usr/bin/env python3

import numpy
import math
import time
import json

# Import the LSM303 module.
import Adafruit_LSM303

calibration = {
    "accelerometer": {
        "min": [-959, -999, -1021],
        "max": [991, 1029, 1153]
    },
    "magnetometer": {
        "min": [-403, -468, -470],
        "max": [662, 395, 575]
    }
}


def calibrate():
    lsm303 = Adafruit_LSM303.LSM303()
    accel_min = [9999, 9999, 9999]
    accel_max = [-9999, -9999, -9999]
    mag_min = [9999, 9999, 9999]
    mag_max = [-9999, -9999, -9999]

    while True:
        accel, mag = lsm303.read()
        for index, item in enumerate(accel):
            accel_max[index] = max(accel_max[index], item)
            accel_min[index] = min(accel_min[index], item)
        for index, item in enumerate(mag):
            mag_max[index] = max(mag_max[index], item)
            mag_min[index] = min(mag_min[index], item)
        calibration = {
            "accel": {
                "min": accel_min,
                "max": accel_max
            },
            "mag": {
                "min": mag_min,
                "max": mag_max
            }
        }
        print(json.dumps(calibration))
        time.sleep(0.5)


def projection_heading(accel, mag):
    """Uses projection on to normal plane to adjust vector.
    """
    (acl_x, acl_y, acl_z) = accel
    (mag_x, mag_y, mag_z) = mag
    mag_x = (mag_x - calibration['magnetometer']['min'][0]) / (calibration[
        'magnetometer']['max'][0] - calibration['magnetometer'][
            'min'][0]) * 2 - 1
    mag_y = (mag_y - calibration['magnetometer']['min'][1]) / (calibration[
        'magnetometer']['max'][1] - calibration['magnetometer'][
            'min'][1]) * 2 - 1
    mag_z = (mag_z - calibration['magnetometer']['min'][2]) / (calibration[
        'magnetometer']['max'][2] - calibration['magnetometer'][
            'min'][2]) * 2 - 1
    a = [acl_x, acl_y, acl_z]
    m = [mag_x, mag_y, mag_z]
    coef = numpy.dot(a, m) / numpy.dot(a, a)
    dir = numpy.subtract(m, numpy.multiply(coef, a))
    (x, y, z) = dir
    h = (math.atan2(y, x) * 180) / math.pi
    if h < 0:
        h += 360
    return h


def trig_heading(accel, mag):
    """Uses pitch and roll to adjust vector.
    """
    (acl_x, acl_y, acl_z) = accel
    (mag_x, mag_y, mag_z) = mag
    mag_x = (mag_x - calibration['magnetometer']['min'][0]) / (calibration[
        'magnetometer']['max'][0] - calibration['magnetometer'][
            'min'][0]) * 2 - 1
    mag_y = (mag_y - calibration['magnetometer']['min'][1]) / (calibration[
        'magnetometer']['max'][1] - calibration['magnetometer'][
            'min'][1]) * 2 - 1
    mag_z = (mag_z - calibration['magnetometer']['min'][2]) / (calibration[
        'magnetometer']['max'][2] - calibration['magnetometer'][
            'min'][2]) * 2 - 1

    # normalize between 0 and 1
    acl_x_norm = acl_x / \
        math.sqrt(acl_x * acl_x + acl_y * acl_y + acl_z * acl_z)
    acl_y_norm = acl_y / \
        math.sqrt(acl_x * acl_x + acl_y * acl_y + acl_z * acl_z)

    pitch = math.asin(-acl_x_norm)
    roll = math.asin(acl_y_norm / math.cos(pitch))
    mag_x_comp = mag_x * math.cos(pitch) + mag_z * math.sin(pitch)
    mag_y_comp = mag_x * math.sin(roll) * math.sin(pitch) + mag_y * \
        math.cos(roll) - mag_z * math.sin(roll) * math.cos(pitch)

    h = (math.atan2(mag_y_comp, mag_x_comp) * 180) / math.pi
    if h < 0:
        h += 360
    return h


def mag_heading(mag):
    (x, y, z) = mag
    x_offset = x - ((calibration['magnetometer']['min'][0] +
                     calibration['magnetometer']['max'][0]) / 2)
    y_offset = y - ((calibration['magnetometer']['min'][1] +
                     calibration['magnetometer']['max'][1]) / 2)

    h = (math.atan2(y_offset, x_offset) * 180) / math.pi
    if h < 0:
        h += 360
    return h

"""
def test():
    cases = {
        "North": {
            "accel": (-22, 40, 994),
            "mag": (430, 199, -671),
            "expected": 0},
        "East": {
            "accel": (-4, 29, 995),
            "mag": (214, 477, -685),
            "expected": 90},
        "South": {
            "accel": (2, 41, 996),
            "mag": (-124, 190, -685),
            "expected": 180},
        "West": {
            "accel": (-10, 55, 998),
            "mag": (164, -72, -656),
            "expected": 270},
    }

    for t, params in cases.items():
        # actual = heading2(params['accel'], params['mag'])
        actual = mag_heading(params['mag'])
        print("Got: ", actual, "\twanted: ", params[
              'expected'], "\tdiff: ", abs(params['expected'] - actual))
"""


def print_heading():
    lsm303 = Adafruit_LSM303.LSM303()
    while True:
        accel, mag = lsm303.read()
        h = trig_heading(accel, mag)
        print(h)
        time.sleep(0.5)

if __name__ == "__main__":
    print_heading()
