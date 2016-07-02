#!/usr/bin/env python3

import os
import time

os.environ["QUICK2WIRE_API_HOME"] = "/usr/local/autonoma/manual/quick2wire-python-api"
os.environ["PYTHONPATH"] += ":/usr/local/autonoma/manual/quick2wire-python-api"

from i2clibraries import i2c_hmc5883l

def init():
    hmc5883l = i2c_hmc5883l.i2c_hmc5883l(1)
    hmc5883l.setContinuousMode()
    hmc5883l.setDeclination(13, 30)
    return hmc5883l

def degrees():
    hmc5883l = init()
    (h, m) = hmc5883l.getHeading()
    return h

if __name__ == "__main__":
    hmc5883l = init()
    while True:
        print(hmc5883l)
        time.sleep(1.0)