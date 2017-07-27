import unittest

from emulation.sensor import SensorEmulator


class TestSensorEmulator(unittest.TestCase):

    def test_distance(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 7.0,
        }
        e = SensorEmulator(sensor, {})
        e.rotate(45.0)
        SensorEmulator.occupied = set({(3, 4)})
        r = e.read()
        self.assertEqual(r, 5)

    def test_quadrant_one(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 3.0,
        }
        e = SensorEmulator(sensor, {})
        e.rotate(45.0)
        SensorEmulator.occupied = set({(1, 1)})
        self.assertEqual(format(e.read(), ".2f"), "1.41")

    def test_quadrant_two(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 45.0,  # degrees
            "maxDistance": 3.0,
        }
        e = SensorEmulator(sensor, {})
        e.rotate(315.0)
        SensorEmulator.occupied = set({(-1, 1)})
        self.assertEqual(format(e.read(), ".2f"), "1.41")

    def test_quadrant_three(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 3.0,
        }
        e = SensorEmulator(sensor, {})
        e.rotate(225.0)
        SensorEmulator.occupied = set({(-1, -1)})
        self.assertEqual(format(e.read(), ".2f"), "1.41")

    def test_quadrant_four(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 3.0,
        }
        e = SensorEmulator(sensor, {})
        e.rotate(115.0)
        SensorEmulator.occupied = set({(1, -1)})
        self.assertEqual(format(e.read(), ".2f"), "1.41")
