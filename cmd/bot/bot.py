#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
from lib import orientation
from lib.servo import Servo
from lib.move import Move
from lib.ir import IR
from lib.ultrasonic import UltraSonic


from util.getch import *


class Bot:

    def __init__(self):
        gpio.setmode(gpio.BOARD)
        self.move = Move()
        self.servo = Servo()
        self.ir = IR()
        self.ultrasonic = UltraSonic()

    def show_hud(self):
        usd = self.ultrasonic.distance()
        if not usd:
            usd = 0
        p = 'Heading={0:0.2f}°\t' + \
            'Speed={1}%\t' + \
            'Servo={2}°\t' + \
            'IR={3:0.2f}cm\t' + \
            'UltraSonic={4:0.2f}cm'
        print(p.format(
            self.orientation.heading(),
            speed,
            self.servo.position(),
            self.ir.distance(),
            usd))

    def manual_control(self):
        print("Use w,a,s,d to move the vehicle. to exit")
        t = 0.2
        speed = 50
        step = 10
        while True:
            self.show_hud()
            k = getch()
            if k == "w":
                self.move.forward(t, speed)
            elif k == "s":
                self.move.backward(t, speed)
            elif k == "a":
                self.move.counter_clockwise(t, speed)
            elif k == "d":
                self.move.clockwise(t, speed)
            elif k == 'r':
                speed = min(100, speed + step)
            elif k == 'f':
                speed = max(0, speed - step)
            elif k == "q":
                self.servo.move(max(-90, self.servo.position() - 45))
            elif k == "e":
                self.servo.move(min(90, self.servo.position() + 45))
            elif k == 'p':
                continue
            elif k == "x":
                break
        print("done")
    gpio.cleanup()

if __name__ == "__main__":
    bot = Bot()
    bot.manual_control()
