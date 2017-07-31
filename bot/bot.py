#!/usr/bin/env python3

import logging
import json
import socket
import time

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
    conn_buffer_size = 4096

    def __init__(self, sensor_constructors, motor_constructors):
        logger.info("Initializing Bot.")
        self._register_sensors(sensor_constructors)
        self._register_motors(motor_constructors)
        self._set_location(0, 0)
        self._set_heading(0)
        self.conn = None

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
            logger.info("Registered" + str(item))

    def _set_location(self, x, y):
        self.location_x = x
        self.location_y = y

    def _set_heading(self, heading):
        self.heading = heading

    def _move(self, distance, destination):
        logger.info("Moving to (%d, %d)" % (destination[0], destination[1]))
        if 'driver' not in self.motors:
            logger.error("No driver loaded, can not move.")
            return
        self.motors['driver'].move(distance, destination)
        self._set_location(destination[0], destination[1])
        self._report('LOCATION{"x":%d, "y":%d}' %
                     (self.location_x, self.location_y))

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
        if sensor_id == "all":
            for sensor, sensor_obj in self.sensors.items():
                self._read(sensor)
        elif sensor_id is None:
            logger.error("sensor_id can not be None")
        elif sensor_id not in self.sensors:
            logger.error("Invalid sensor: %s" % sensor_id)
        else:
            r = self.sensors[sensor_id].read()
            self._report(self._create_payload(sensor_id, r))

    def _create_payload(self, sensor_id, reading):
        heading = self.heading
        if sensor_id == 'irdistance':
            heading = (self.heading + self.servo_position) % 360
        payload = {
            'sensorId': sensor_id,
            'output': reading,
            'heading': heading,
            'x': int(self.location_x),
            'y': int(self.location_y)
        }
        if sensor_id != 'compass' and reading != 0 and reading != self.sensors[sensor_id].metadata['maxDistance']:
            logger.info("Occupied dist" + str(payload))
        return "READING" + json.dumps(payload)

    def _report(self, payload):
        if self.conn:
            logger.debug("Sending " + payload)
            self.conn.sendall((payload + '\n').encode())
        else:
            logging.error("Not connected. Can not send.")

    def _report_sensors(self):
        for senor_id, sensor in self.sensors.items():
            self._report("SENSOR" + json.dumps(sensor.metadata))
            time.sleep(0.01)

    def _handle(self, payload):
        logger.debug(payload)
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
                logger.info('Connection address: ' + addr[0])
                self._report_sensors()
                smsg = ""
                while True:
                    buffer = self.conn.recv(self.conn_buffer_size)
                    smsg += str(buffer, "utf-8")
                    logger.debug("Received " + str(buffer, "utf-8"))
                    if not buffer:
                        break
                    cmds = smsg.split('\n')
                    if smsg.endswith('\n'):
                        smsg = ''
                    else:
                        smsg = cmds[-1]
                        cmds = cmds[:-1]
                    for cmd in cmds:
                        if cmd:
                            self._handle(cmd)
            except KeyboardInterrupt:
                logger.info("User exit.")
                exit = True
                break
            if exit:
                break
        if self.conn:
            self.conn.close()
            self.conn = None
            logger.info("Connection closed.")

if __name__ == "__main__":
    print("Run Bernie via './bernie.py' or the emulator via './emulator.py'")
