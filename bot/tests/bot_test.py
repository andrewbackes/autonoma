import unittest

from bot.emulate import *


class TestBot(unittest.TestCase):

    def test_syntax(self):
        bot = new_emulator()
