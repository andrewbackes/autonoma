#!/usr/bin/env python3

# from pi:
# red, brown, black, white
# +, SDA, SLC, -

# from sensor:
# red,  orange, yellow, green,  blue,   black
# 5v,   enable, mode,   scl,    sda,    ground

import Adafruit_GPIO.I2C as I2C
import time

##########################################################################
#
#  Garmin LIDAR-Lite V3 range finder
#
##########################################################################


class GLL:
    i2c = None

    __GLL_ACQ_COMMAND = 0x00
    __GLL_STATUS = 0x01
    __GLL_SIG_COUNT_VAL = 0x02
    __GLL_ACQ_CONFIG_REG = 0x04
    __GLL_VELOCITY = 0x09
    __GLL_PEAK_CORR = 0x0C
    __GLL_NOISE_PEAK = 0x0D
    __GLL_SIGNAL_STRENGTH = 0x0E
    __GLL_FULL_DELAY_HIGH = 0x0F
    __GLL_FULL_DELAY_LOW = 0x10
    __GLL_OUTER_LOOP_COUNT = 0x11
    __GLL_REF_COUNT_VAL = 0x12
    __GLL_LAST_DELAY_HIGH = 0x14
    __GLL_LAST_DELAY_LOW = 0x15
    __GLL_UNIT_ID_HIGH = 0x16
    __GLL_UNIT_ID_LOW = 0x17
    __GLL_I2C_ID_HIGHT = 0x18
    __GLL_I2C_ID_LOW = 0x19
    __GLL_I2C_SEC_ADDR = 0x1A
    __GLL_THRESHOLD_BYPASS = 0x1C
    __GLL_I2C_CONFIG = 0x1E
    __GLL_COMMAND = 0x40
    __GLL_MEASURE_DELAY = 0x45
    __GLL_PEAK_BCK = 0x4C
    __GLL_CORR_DATA = 0x52
    __GLL_CORR_DATA_SIGN = 0x53
    __GLL_ACQ_SETTINGS = 0x5D
    __GLL_POWER_CONTROL = 0x65

    def __init__(self, address=0x62, rate=10):
        # self.i2c = I2C(address)
        self.i2c = I2C.get_i2c_device(address)
        self.rate = rate

        # ----------------------------------------------------------------------
        # Set to continuous sampling after initial read.
        # ----------------------------------------------------------------------
        self.i2c.write8(self.__GLL_OUTER_LOOP_COUNT, 0xFF)

        # ----------------------------------------------------------------------
        # Set the sampling frequency as 2000 / Hz:
        # 10Hz = 0xc8
        # 20Hz = 0x64
        # 100Hz = 0x14
        # ----------------------------------------------------------------------
        self.i2c.write8(self.__GLL_MEASURE_DELAY, int(2000 / rate))

        # ----------------------------------------------------------------------
        # Include receiver bias correction 0x04
        # ----------------------------------------------------------------------
        self.i2c.write8(self.__GLL_ACQ_COMMAND, 0x04)

        # ----------------------------------------------------------------------
        # Acquisition config register:
        # 0x01 Data ready interrupt
        # 0x20 Take sampling rate from MEASURE_DELAY
        # ----------------------------------------------------------------------
        self.i2c.write8(self.__GLL_ACQ_CONFIG_REG, 0x21)

    def read(self):
        # ----------------------------------------------------------------------
        # Distance is in cm
        # Velocity is in cm between consecutive reads; sampling rate converts these to a velocity
        # Reading the list from 0x8F seems to get the previous reading, probably cached for the sake
        # of calculating the velocity next time round.
        # ----------------------------------------------------------------------
        '''
        gll_bytes = self.i2c.readList(0x80 | self.__GLL_LAST_DELAY_HIGH, 2)
        dist1 = gll_bytes[0]
        dist2 = gll_bytes[1]
        distance = ((dist1 << 8) + dist2) / 100
        '''

        dist1 = self.i2c.readU8(self.__GLL_FULL_DELAY_HIGH)
        dist2 = self.i2c.readU8(self.__GLL_FULL_DELAY_LOW)
        distance = ((dist1 << 8) + dist2) / 100

        velocity = -self.i2c.readS8(self.__GLL_VELOCITY) * self.rate / 100
        return distance, velocity

if __name__ == "__main__":
    lidar = GLL()
    while True:
        print(lidar.read())
        time.sleep(0.5)
