#!/usr/bin/env python3

import logging
import json
import socket

from sensors.ultrasonic import Ultrasonic
from sensors.compass import Compass
from sensors.irdist import IRDistance
from sensors.irprox import IRProximity

from sensors.config import settings as sensor_settings

from motors.driver import Driver
from motors.servo import Servo

from motors.config import settings as motor_settings


logger = logging.getLogger(__name__)


class Bot(object):

    bind_ip = '0.0.0.0'
    bind_port = 9091
    conn_buffer_size = 256

    def __init__(self, sensor_constructors, motor_constructors):
        logger.info("Initializing Bot.")
        self._register_sensors(sensor_constructors)
        self._register_motors(motor_constructors)
        self._set_location(0, 0)

    def _register_sensors(self, sensor_constructors):
        self.sensors = {}
        self._register(
            sensor_constructors,
            sensor_settings.items(),
            self.sensors)

    def _register_motors(self, motor_constructors):
        self.motors = {}
        self._register(
            motor_constructors,
            motor_settings.items(),
            self.motors)

    def _register(self, constructors, items, reference_map):
        for id, item in items:
            reference_map[id] = constructors[item['type']](
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
        self.motors['driver'].rotate(heading)
        self._set_heading(heading)

    def _look(self, degrees):
        if 'servo' not in self.motors:
            logger.error("No servo loaded, can not look around.")
            return
        self.motors['servo'].position(degrees)
        self.servo_position = degrees

    def _read(self, sensor_id):
        if sensor_id is None:
            logger.error("Can not read sensor_id: 'None'")
            return
        if sensor_id == "all":
            for sensor in self.sensors.items():
                self._read(sensor)
        r = self.sensors[sensor_id].read()

    def _send(self, payload):
        if self.conn:
            self.conn.sendto(payload)
        else:
            logging.error("Not connected. Can not send.")

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
            try:
                self.conn, addr = s.accept()
                print('Connection address:', addr)
                while True:
                    msg = self.conn.recv(self.conn_buffer_size)
                    if not msg:
                        break
                    self._handle(str(msg, "utf-8"))
            except KeyboardInterrupt:
                print("User exit.")
                exit = True
                break
            self.conn.close()
            self.conn = None
            print("Connection closed.")
            if exit:
                break

if __name__ == "__main__":
    print("Run Bernie via './bernie.py' or the emulator via './emulator.py'")
