import unittest

from bot.emulation.emulator import Emulator


class TestEmulator(unittest.TestCase):

    def test_distance(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 5.0,
        }
        e = Emulator(sensor, {})
        e.face(45.0)
        Emulator.occupied = set({(2, 2)})
        r = e.read()
        self.assertGreater(r, 2.8)
        self.assertLess(r, 2.9)
