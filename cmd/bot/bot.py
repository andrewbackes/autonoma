#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
from lib.orientation import Orientation
from lib.move import Move
from lib.roofmount import RoofMount
from lib.ir import IR
from lib.ultrasonic import UltraSonic

from util.getch import *
from util.tcp import TCP

import json
import sys
import os
import time


class Bot:
    _config = {
        "hbridge": {"enabled": False},
        "roofmount": {"enabled": False},
        "orientation": {"enabled": False},
        "ultrasonic": {"enabled": False},
        "ir": {"enabled": False},
    }
    _sensor_readers = {}
    _move = None
    _roofmount = None

    def __init__(self, config):
        gpio.setmode(gpio.BOARD)
        self._config.update(config)
        # controls:
        if self._config['hbridge'] and self._config['hbridge']['enabled']:
            self._move = Move()
        if self._config['roofmount'] and self._config['roofmount']['enabled']:
            self._roofmount = RoofMount(self._config['roofmount'])
            self._sensor_readers['roofmount'] = self._roofmount.get_readings

        # sensors readers:
        if self._config['orientation'] \
                and self._config['orientation']['enabled']:
            orientation = Orientation()
            self._sensor_readers['orientation'] = orientation.heading
        if self._config['ir'] and self._config['ir']['enabled']:
            ir = IR()
            self._sensor_readers['ir'] = ir.distance
        if self._config['ultrasonic'] \
                and self._config['ultrasonic']['enabled']:
            ultrasonic = UltraSonic()
            self._sensor_readers['ultrasonic'] = ultrasonic.distance

    def __gpio_reset(self):
        gpio.cleanup()
        if self._config['hbridge'] and self._config['hbridge']['enabled']:
            self._move = Move()
        if self._config['servo'] and self._config['servo']['enabled']:
            self._servo = Servo(self._config['servo'])
        if self._config['ultrasonic'] \
                and self._config['ultrasonic']['enabled']:
            ultrasonic = UltraSonic()
            self._sensor_readers['ultrasonic'] = ultrasonic.distance

    def __del__(self):
        print("done")
        gpio.cleanup()

    def get_readings(self):
        r = {}
        if self._roofmount is not None:
            r.update(self._roofmount.get_readings())
        for name, read in self._sensor_readers.items():
            r[name] = read()
        r['timestamp'] = time.time()
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
            cmd = None
            if k == "w":
                cmd = {'command': 'move', 'direction': 'forward',
                       'speed': speed, 'time': t}
            elif k == "s":
                cmd = {'command': 'move', 'direction': 'backward',
                       'speed': speed, 'time': t}
            elif k == "a":
                cmd = {'command': 'move', 'direction': 'counter_clockwise',
                       'speed': speed, 'time': t}
            elif k == "d":
                cmd = {'command': 'move', 'direction': 'clockwise',
                       'speed': speed, 'time': t}
            elif k == 't':
                speed = min(100, speed + step)
            elif k == 'g':
                speed = max(0, speed - step)
            elif k == "q" and self._roofmount is not None:
                cmd = {'command': 'horizontal_position',
                       'position': (self._roofmount.horizontal_position() +
                                    (360 - 15)) % 360}
            elif k == "e" and self._roofmount is not None:
                cmd = {'command': 'horizontal_position',
                       'position': (self._roofmount.horizontal_position() +
                                    15) % 360}
            elif k == "r" and self._roofmount is not None:
                cmd = {'command': 'vertical_position',
                       'position': self._roofmount.vertical_position() + 15}
            elif k == "f" and self._roofmount is not None:
                cmd = {'command': 'vertical_position',
                       'position': self._roofmount.vertical_position() - 15}
            elif k == 'p':
                continue
            elif k == "x":
                break
            if cmd is not None:
                self.__execute(cmd)

    def network_control(self):
        self.tcp = TCP()
        self.tcp.listen(self.__handler)

    def __handler(self, payload):
        print('Handling ' + payload)
        try:
            cmd = json.loads(payload)
        except:
            print("Could not decode json")
            return
        self.__execute(cmd)

    def __execute(self, cmd):
        # Drive controls:
        if cmd['command'] == 'move' and cmd['direction'] == 'forward':
            if self._move:
                self._move.forward(cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' and cmd['direction'] == 'backward':
            if self._move:
                self._move.backward(cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' \
                and cmd['direction'] == 'counter_clockwise':
            if self._move:
                self._move.counter_clockwise(
                    cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' and cmd['direction'] == 'clockwise':
            if self._move:
                self._move.clockwise(cmd['time'], cmd['speed'])

        # View controls:
        elif cmd['command'] == 'horizontal_position':
            if self._roofmount:
                self._roofmount.set_horizontal_position(cmd['position'])
        elif cmd['command'] == 'vertical_position':
            if self._roofmount:
                self._roofmount.set_vertical_position(cmd['position'])

        # Sensor controls:
        elif cmd['command'] == 'get_readings':
            if self._servo:
                self.tcp.send(self.get_readings())

        # Communication controls:
        elif cmd['command'] == 'isready':
            self.tcp.send('{"status":"readyok"}')
        elif cmd['command'] == 'reset':
            self.__gpio_reset()


if __name__ == "__main__":
    # check args:
    if len(sys.argv) <= 1:
        print("Please specify --network or --manual")
        sys.exit(1)

    # set working dir:
    abspath = os.path.abspath(__file__)
    dname = os.path.dirname(abspath)
    os.chdir(dname)

    # Control bot:
    f = open('config.json')
    settings = json.load(f)
    f.close()
    bot = Bot(settings)
    if sys.argv[1] == '--network':
        bot.network_control()
    elif sys.argv[1] == '--manual':
        bot.manual_control()
    elif sys.argv[1] == '--test':
        print("todo")
    else:
        print("unknown arguement")
    del(bot)
