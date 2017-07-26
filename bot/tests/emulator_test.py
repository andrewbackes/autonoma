import unittest

from bot.emulation.emulator import Emulator


class TestEmulator(unittest.TestCase):

    def test_distance(self):
        sensor = {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "maxDistance": 7.0,
        }
        e = Emulator(sensor, {})
        e.face(45.0)
        Emulator.occupied = set({(3, 4)})
        r = e.read()
        self.assertEqual(r, 5)
