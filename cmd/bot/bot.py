#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
from lib.orientation import Orientation
from lib.servo import Servo
from lib.move import Move
from lib.ir import IR
from lib.ultrasonic import UltraSonic
from lib.lidar import Lidar

from util.getch import *
from util.tcp import TCP

import json
import sys
import os


class Bot:
    _config = {}
    _sensor_readers = {}

    def __init__(self, config):
        gpio.setmode(gpio.BOARD)
        self._config.update(config)
        # controls:
        if self._config['hbridge'] and self._config['hbridge']['enabled']:
            self.move = Move()
        if self._config['servo'] and self._config['servo']['enabled']:
            self.servo = Servo(self._config['servo'])

        # sensors:
        if self._config['orientation'] and self._config['orientation']['enabled']:
            orientation = Orientation()
            self._sensor_readers['orientation'] = orientation.heading
        if self._config['ir'] and self._config['ir']['enabled']:
            ir = IR()
            self._sensor_readers['ir'] = ir.distance
        if self._config['ultrasonic'] and self._config['ultrasonic']['enabled']:
            ultrasonic = UltraSonic()
            self._sensor_readers['ultrasonic'] = ultrasonic.distance
        if self._config['lidar'] and self._config['lidar']['enabled']:
            lidar = Lidar()
            self._sensor_readers['lidar'] = lidar.distance

    def __del__(self):
        print("done")
        gpio.cleanup()

    def get_readings(self):
        r = {}
        if self.servo:
            r['servo'] = self.servo.position()
        for name, read in self._sensor_readers.items():
            r[name] = read()
        r['timestamp'] = time.time()
        return json.dumps(r)

    def manual_control(self):
        print("Use w,a,s,d to move the vehicle. to exit")
        t = 0.2
        speed = 50
        step = 10
        servo_step = 30
        while True:
            print(self.get_readings())
            print('Speed ', speed)
            k = getch()
            cmd = {}
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
            elif k == 'r':
                speed = min(100, speed + step)
            elif k == 'f':
                speed = max(0, speed - step)
            elif k == "q":
                cmd = {'command': 'servo',
                       'position': max(-180, self.servo.position() - servo_step)}
            elif k == "e":
                cmd = {'command': 'servo',
                       'position': min(180, self.servo.position() + servo_step)}
            elif k == 'p':
                continue
            elif k == "x":
                break
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
        if cmd['command'] == 'move' and cmd['direction'] == 'forward':
            if self.move:
                self.move.forward(cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' and cmd['direction'] == 'backward':
            if self.move:
                self.move.backward(cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' and cmd['direction'] == 'counter_clockwise':
            if self.move:
                self.move.counter_clockwise(
                    cmd['time'], cmd['speed'])
        elif cmd['command'] == 'move' and cmd['direction'] == 'clockwise':
            if self.move:
                self.move.clockwise(cmd['time'], cmd['speed'])
        elif cmd['command'] == 'servo':
            if self.servo:
                self.servo.move(cmd['position'])
        elif cmd['command'] == 'get_readings':
            if self.servo:
                self.tcp.send(self.get_readings())
        elif cmd['command'] == 'isready':
            self.tcp.send('{"status":"readyok"}')


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
