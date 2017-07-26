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
        r = format(e.read(), '.2f')
        self.assertEqual(r, '2.83')
