import unittest
import logging


from emulation.courses import box
from emulator import Emulator

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)


class TestBot(unittest.TestCase):

    def setUp(self):
        self.bot = Emulator(box.occupied)

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
        test_payload = """
            {  "action": "read",
               "sensor_id": "all" }
        """
        self.bot._handle(test_payload)

        def capture_msg(msg):
            logger.debug(msg)
        self.bot._report = capture_msg

        self.assertTrue(False)
