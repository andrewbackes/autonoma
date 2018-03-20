#!/usr/bin/env python3

import logging
import socket

logger = logging.getLogger(__name__)


class TCP:
    bind_ip = '0.0.0.0'
    bind_port = 9091
    conn_buffer_size = 4096

    def __init__(self):
        self.conn = None

    def send(self, msg):
        print('Sending ' + msg)

    def listen(self, handler):
        logger.info("Listening for TCP/IP connections.")
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.bind((self.bind_ip, self.bind_port))
        s.listen(1)
        exit = False
        while True:
            try:
                self.conn, addr = s.accept()
                logger.info('Connection address: ' + addr[0])
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
                            handler(cmd)
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
