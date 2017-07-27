#!/usr/bin/env python3

import logging
import json
import socket

from bot.sensors.ultrasonic import Ultrasonic
from bot.sensors.compass import Compass
from bot.sensors.irdist import IRDistance
from bot.sensors.irprox import IRProximity

from bot.sensors.config import settings as sensor_settings

from bot.motors.driver import Driver
from bot.motors.servo import Servo

from bot.motors.config import settings as motor_settings


logger = logging.getLogger(__name__)


class Bot(object):

    bind_ip = '0.0.0.0'
    bind_port = 9091
    conn_buffer_size = 256

    def __init__(self, sensor_factory, motor_factory):
        logger.info("Initializing Bot.")
        self._register_sensors(sensor_factory)
        self._register_motors(motor_factory)
        self._set_location(0, 0)

    def _register_sensors(self, sensor_factory):
        default = {
            "ultrasonic": Ultrasonic,
            "irdistance": IRDistance,
            "irproximity": IRProximity,
            "compass": Compass,
        }
        self.sensors = {}
        self._register(
            sensor_factory,
            default,
            sensor_settings.items(),
            self.sensors)

    def _register_motors(self, motor_factory):
        default = {
            "driver": Driver,
            "servo": Servo,
        }
        self.motors = {}
        self._register(
            motor_factory,
            default,
            motor_settings.items(),
            self.motors)

    def _register(self, factory, default_factory, items, member_map):
        member_factory = factory if factory else default_factory
        for id, item in items:
            member_map[id] = member_factory[item['type']](
                item['metadata'], item['config'])
            logger.info("Registered", item)

    def _set_location(self, x, y):
        self.location_x = x
        self.location_y = y

    def _set_heading(self, heading):
        self.heading = heading

    def _move(self, distance, destination):
        if 'driver' not in self.motors:
            logger.error("No driver loaded, can not move.")
            return
        self.motors['driver'].move(distance, destination)
        self._set_location(destination[0], destination[1])

    def _rotate(self, heading):
        if 'driver' not in self.motors:
            logger.error("No driver loaded, can not rotate.")
            return
        self.motors['driver'].rotate(distance)
        self._set_heading(heading)

    def _rotate(self, heading):
        if 'servo' not in self.motors:
            logger.error("No servo loaded, can not look around.")
            return
        self.motors['servo'].position(degrees)
        self.servo_position = degrees

    def _read(self, sensor_id):
        if sensor is None:
            logger.error("Can not read sensor 'None'")
            return
        if sensor == "all":
            for sensor in self.sensors.items():
                self._read(sensor)
        r = self.sensors[sensor_id].read()

    def _send(self, payload):
        if self.conn:
            self.conn.sendto(payload)

    def _handle(self, payload):
        cmd = json.loads(payload)
        try:
            if cmd['action'] == "move":
                """ Expecting:
                {   "action": "move",
                    "distance": 10,
                    "destination": {
                        "x": 3,
                        "y": 4 }}
                """
                self._move(cmd['distance'],
                           (cmd['destination']['x'], cmd['destination']['y']))

            elif cmd['action'] == "rotate":
                """ Expecting:
                {  "action": "rotate",
                   "heading": 45.0 }
                """
                self._rotate(cmd['heading'])

            elif cmd['action'] == "look":
                """ Expecting:
                {  "action": "look",
                   "degrees": -15.0 }
                """
                self._look(cmd['degrees'])

            elif cmd['action'] == "read":
                """ Expecting:
                {  "action": "read",
                   "sensor_id": "all" }
                """
                self._read(cmd['sensor_id'])

            else:
                logger.error("Unknown action: %d" % cmd['action'])

        except KeyError:
            logger.error("Missing key from payload: %s" % payload)

    def start(self):
        logger.info("Bot started.")
        self._await_tcp()
        logger.info("Bot stopped.")

    def _await_tcp(self):
        logger.info("Listening for TCP/IP connections.")
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.bind((self.bind_ip, self.bind_port))
        s.listen(1)
        exit = False
        while True:
            self.conn, addr = s.accept()
            print('Connection address:', addr)
            while True:
                try:
                    msg = self.conn.recv(self.conn_buffer_size)
                    if not msg:
                        break
                    self._handle(str(msg, "utf-8"))
                except KeyboardInterrupt:
                    print("Exit")
                    exit = True
                    break
            self.conn.close()
            self.conn = None
            print("Connection closed.")
            if exit:
                break
