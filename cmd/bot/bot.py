#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
from lib.orientation import Orientation
from lib.servo import Servo
from lib.move import Move
from lib.ir import IR
from lib.ultrasonic import UltraSonic


from util.getch import *
from util.tcp import TCP

import json
import sys


class Bot:

    def __init__(self):
        gpio.setmode(gpio.BOARD)
        self.move = Move()
        self.orientation = Orientation()
        self.servo = Servo()
        self.ir = IR()
        self.ultrasonic = UltraSonic()

    def __del__(self):
        print("done")
        gpio.cleanup()

    def get_readings(self):
        usd = self.ultrasonic.distance()
        if not usd:
            usd = 0
        r = {
            'heading': self.orientation.heading(),
            'servo': self.servo.position(),
            'ir': self.ir.distance(),
            'ultrasonic': usd
        }
        return json.dumps(r)

    def manual_control(self):
        print("Use w,a,s,d to move the vehicle. to exit")
        t = 0.2
        speed = 50
        step = 10
        while True:
            print(self.get_readings())
            print('Speed ', speed)
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

    def __handler(self, payload):
        print('Handling ' + payload)
        try:
            p = json.loads(payload)
        except:
            print("Could not decode json")
            return
        if p['command'] == 'move' and p['direction'] == 'forward':
            self.move.forward(p['time'], p['speed'])
        elif p['command'] == 'move' and p['direction'] == 'backward':
            self.move.backward(p['time'], p['speed'])
        elif p['command'] == 'move' and p['direction'] == 'counter_clockwise':
            self.move.counter_clockwise(p['time'], p['speed'])
        elif p['command'] == 'move' and p['direction'] == 'clockwise':
            self.move.clockwise(p['time'], p['speed'])
        elif p['command'] == 'servo':
            self.servo.move(p['position'])
        elif p['command'] == 'get_readings':
            self.tcp.send(json.dumps(self.get_readings()))

    def network_control(self):
        self.tcp = TCP()
        self.tcp.listen(self.__handler)


if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print("Please specify --network or --manual")
        sys.exit(1)
    bot = Bot()
    if sys.argv[1] == '--network':
        bot.network_control()
    elif sys.argv[1] == '--manual':
        bot.manual_control()
    else:
        print("unknown arguement")
    del(bot)
