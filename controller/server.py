#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#import RPi.GPIO as gpio
import socket

'''
import move

def handle_move(dir, t, power):
    if k == "forward":
        move.forward(t, power)
    elif k == "backward":
        move.backward(t, power)
    elif k == "counter_clockwise":
        move.counter_clockwise(t, power)
    elif k == "clockwise":
        move.clockwise(t, power)
'''
def handle(data):
    # ex: move forward 1000 80
    terms  = data.split(" ")
    if terms[0] == "move" and len(terms) >= 4:
        pass
        #handle_move(terms[1], terms[2], terms[3])
    print(terms)

def listen_and_serve():    
    print("Listening for connection.")
    TCP_IP = '127.0.0.1'
    TCP_PORT = 9091
    BUFFER_SIZE = 256
    
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind((TCP_IP, TCP_PORT))
    s.listen(1)

    while True:
        conn, addr = s.accept()
        print('Connection address:', addr)
        while True:
            msg = conn.recv(BUFFER_SIZE)
            if not msg: break
            handle(msg)
        conn.close()
        print("Connection closed.")


if __name__ == "__main__":
    try:
        listen_and_serve()
    except KeyboardInterrupt:
        print("Exit")