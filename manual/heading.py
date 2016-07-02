#!/usr/bin/env python3

import os

#export QUICK2WIRE_API_HOME=/usr/local/autonoma/manual/quick2wire-python-api
#export PYTHONPATH=$PYTHONPATH:$QUICK2WIRE_API_HOME
os.environ["QUICK2WIRE_API_HOME"] = "/usr/local/autonoma/manual/quick2wire-python-api"
os.environ["PYTHONPATH"] += ":/usr/local/autonoma/manual/quick2wire-python-api"

from i2clibraries import i2c_hmc5883l

def init():
    hmc5883l.setContinuousMode()
    hmc5883l.setDeclination(13, 30)

def degrees():
    init()
    (h, m) = hmc5883l.getHeading()
    return h

if __name__ == "__main__":
    while True:
        print(hmc5883l)
        time.sleep(1.0)