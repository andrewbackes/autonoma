#!/usr/bin/env python3

from bot import Bot

from sensors.ultrasonic import Ultrasonic
from sensors.compass import Compass
from sensors.irdist import IRDistance
from sensors.irprox import IRProximity

from motors.driver import Driver
from motors.servo import Servo


class Bernie(Bot):

    def __init__(self):
        sensors = {
            "ultrasonic": Ultrasonic,
            "irdistance": IRDistance,
            "irproximity": IRProximity,
            "compass": Compass,
        }
        motors = {
            "driver": Driver,
            "servo": Servo,
        }
        super().__init__(sensors, motors)

if __name__ == "__main__":
    print("Running Bernie")
    bernie = Bernie()
    bernie.start()
