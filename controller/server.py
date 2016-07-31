#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import RPi.GPIO as gpio
import socket
import move

def handle_move(dir, t, power):
    if dir == "forward":
        move.forward(t, power)
    elif dir == "backward":
        move.backward(t, power)
    elif dir == "counter_clockwise":
        move.counter_clockwise(t, power)
    elif dir == "clockwise":
        move.clockwise(t, power)

def handle(data):
    # ex: move forward 1000 80
    stripped = data.strip("\n").strip("\r")
    terms  = stripped.split(" ")
    if terms[0] == "move" and len(terms) >= 4:
        pass
        handle_move(terms[1], terms[2], terms[3])
    print(terms)

def listen_and_serve():    
    print("Listening for connection.")
    TCP_IP = '10.0.0.14'
    TCP_PORT = 9091
    BUFFER_SIZE = 256
    
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind((TCP_IP, TCP_PORT))
    s.listen(1)
    exit = False
    while True:
        conn, addr = s.accept()
        print('Connection address:', addr)
        while True:
            try:
                msg = conn.recv(BUFFER_SIZE)
                if not msg: break
                handle(str(msg, "utf-8"))
            except KeyboardInterrupt:
                print("Exit")
                exit = True
                break
        conn.close()
        print("Connection closed.")
        if exit:
            break


if __name__ == "__main__":
    listen_and_serve()