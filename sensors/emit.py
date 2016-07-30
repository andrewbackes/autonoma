#!/usr/bin/env python3

import socket

import ir
import echo
#import heading

ip = "10.0.0.11"
port =  "9090"
interval = 0.01 

def udp_send(msg):
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM) # UDP
    sock.sendto(bytes(msg, "utf-8"), (ip, port))


def send_payload(sensor, location, reading):
    upd_send(sensor + " " + location + " " + reading)


def start():
    print("Emitting sensor data.")
    while True:
        send_payload("ir", "left", ir.blocked('left'))
        send_payload("ir", "right", ir.blocked('right'))
        send_payload("echo", "front", echo.distance(echo.sensors['front']))
        send_payload("echo", "back", echo.distance(echo.sensors['back']))
        #send_payload("heading", "front", heading.degrees())
        time.sleep(interval)


if __name__ == "__main__":
    start()