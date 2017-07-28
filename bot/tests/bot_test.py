import unittest
import logging
import json

from emulation.courses import box
from emulator import Emulator
from emulation.sensor import SensorEmulator

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


class TestBot(unittest.TestCase):

    def setUp(self):
        self.bot = Emulator(box.occupied)
        SensorEmulator.reset()

    def test_handle_move(self):
        '''Test move handler.
        '''
        test_payload = """
            {   "action": "move",
                "distance": 10,
                "destination": {
                    "x": 0,
                    "y": 10 }}
            """
        self.bot._handle(test_payload)
        self.assertEqual(self.bot.location_x, 0)
        self.assertEqual(self.bot.location_y, 10)

    def test_handle_rotate(self):
        '''Test rotation handler.
        '''
        test_payload = """
            {  "action": "rotate",
               "heading": 45.0 }
            """
        self.bot._handle(test_payload)
        self.assertEqual(self.bot.heading, 45.0)

    def test_handle_look(self):
        '''Test look handler.
        '''
        test_payload = """
            {  "action": "look",
               "degrees": -15.0 }
        """
        self.bot._handle(test_payload)
        self.assertEqual(self.bot.servo_position, -15)

    def test_handle_read(self):
        '''Test sensor reader handler.
        '''
        logger.debug(SensorEmulator.heading)
        test_payload = """
            {  "action": "read",
               "sensor_id": "all" }
        """
        self.bot._handle(test_payload)
        sensor_ids = set({
            "front_ultrasonic",
            "read_ultrasonic",
            "left_irdistance",
            "right_irdistance",
            "irproximity",
        })
        captured_sensors = set()

        def capture_msg(payload):
            json.loads(payload)
            captured_sensors.add(payload['sensor_id'])
        self.bot._report = capture_msg
        # self.assertEquals(captured_sensors, sensor_ids)
