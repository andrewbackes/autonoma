#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
from lib import orientation
from lib.servo import Servo
from lib.move import Move
from lib.ir import IR
from lib.ultrasonic import UltraSonic


from util.getch import *


def repl():
    gpio.setmode(gpio.BOARD)
    move = Move()
    servo = Servo()
    ir = IR()
    ultrasonic = UltraSonic()

    print("Use w,a,s,d to move the vehicle. to exit")
    t = 0.2
    speed = 50
    step = 10
    servo_pos = 0
    servo.move(servo_pos)
    while True:
        print('Heading={0:0.2f}°\tSpeed={1}%\tIR={2:0.2f}cm\tUltraSonic={3:0.2f}cm'.format(
            orientation.heading(), speed, ir.distance(), ultrasonic.distance()))

        k = getch()
        if k == "w":
            move.forward(t, speed)
        elif k == "s":
            move.backward(t, speed)
        elif k == "a":
            move.counter_clockwise(t, speed)
        elif k == "d":
            move.clockwise(t, speed)
        elif k == 'r':
            speed = min(100, speed + step)
        elif k == 'f':
            speed = max(0, speed - step)
        elif k == "q":
            servo_pos = max(-90, servo_pos - 45)
            servo.move(servo_pos)
        elif k == "e":
            servo_pos = min(90, servo_pos + 45)
            servo.move(servo_pos)
        elif k == 'p':
            continue
        elif k == "x":
            break
    print("done")
    gpio.cleanup()

if __name__ == "__main__":
    repl()
