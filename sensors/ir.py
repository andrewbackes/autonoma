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
    # https://www.upgradeindustries.com/product/58/Sharp-10-80cm-Infrared-Distance-Sensor-(GP2Y0A21YK0F)
    # http://www.instructables.com/id/Get-started-with-distance-sensors-and-Arduino/
    SPI_PORT   = 0
    SPI_DEVICE = 0
    mcp = Adafruit_MCP3008.MCP3008(spi=SPI.SpiDev(SPI_PORT, SPI_DEVICE))
    value = mcp.read_adc(2)
    cm = 0
    if value:
        distance = 12343.85 * (value**-1.15)
    return cm
