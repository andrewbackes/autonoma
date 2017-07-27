import unittest


from emulation.courses import box
from emulator import Emulator


class TestBot(unittest.TestCase):

    def test_syntax(self):
        '''Test to run through code for syntax errors.
        '''
        bot = Emulator(box.occupied)

    def test_handle_move(self):
        '''Test move handler.
        '''
        pass

    def test_handle_rotate(self):
        '''Test rotation handler.
        '''
        pass

    def test_handle_look(self):
        '''Test look handler.
        '''
        pass

    def test_handle_read(self):
        '''Test sensor reader handler.
        '''
        pass
