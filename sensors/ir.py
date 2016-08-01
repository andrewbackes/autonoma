#!/usr/bin/env python3

import RPi.GPIO as gpio
import time
import Adafruit_GPIO.SPI as SPI
import Adafruit_MCP3008

sensors = {
    'left': 36,
    'right': 32
}

def blocked(sensor="all"):
    gpio.setmode(gpio.BOARD)
    if sensor == 'all':
        gpio.setup(sensors['left'], gpio.IN)
        gpio.setup(sensors['right'], gpio.IN)
        blocked = gpio.input(sensors['left']) == 0 and gpio.input(sensors['right']) == 0
    else:
        gpio.setup(sensors[sensor], gpio.IN)
        blocked = gpio.input(sensors[sensor]) == 0
    gpio.cleanup()
    return blocked

def distance():
    # Hardware SPI configuration:
    SPI_PORT   = 0
    SPI_DEVICE = 0
    mcp = Adafruit_MCP3008.MCP3008(spi=SPI.SpiDev(SPI_PORT, SPI_DEVICE))
    value = mcp.read_adc(0)
    cm = (70/837)*value
    return cm