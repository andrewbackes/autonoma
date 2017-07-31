#!/usr/bin/env python3

settings = {
    # Compass:
    "compass": {
        "metadata": {"id": "compass"},
        "config": {},
        "type": "compass",
    },

    # Ultrasonics:
    "front_ultrasonic": {
        "metadata": {
            "id": "front_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "yOffset": 10.0,  # cm
            "maxDistance": 18.0,  # cm
        },
        "config": {"echo": 16, "trigger": 12},
        "type": "ultrasonic",
    },
    "rear_ultrasonic": {
        "metadata": {
            "id": "rear_ultrasonic",
            "coneWidth": 15.0,  # degrees
            "yOffset": -10.0,  # cm
            "angleOffset": 180.0,  # degrees
            "maxDistance": 18.0,  # cm
        },
        "config": {"echo": 18, "trigger": 22},
        "type": "ultrasonic",
    },

    # IR Proximity:

    # "left_irproximity": {
    #    "metadata": {
    #        "id": "left_irproximity",
    #        "xOffset": -10.0,  # cm
    #        "inclusive": True,
    #        "maxDistance": 10.0,  # cm
    #    },
    #    "config": {"pin": 36},
    #    "type": "irproximity",
    # },
    # "right_irproximity": {
    #    "metadata": {
    #        "id": "right_irproximity",
    #        "xOffset": 10.0,  # cm
    #        "inclusive": True,
    #        "maxDistance": 10.0,  # cm
    #    },
    #    "config": {"pin": 32},
    #    "type": "irproximity",
    # },

    # IR Distance:
    "irdistance": {
        "metadata": {
            "id": "irdistance",
            "maxDistance": 80.0,  # cm
            "minDistance": 10.0,  # cm
        },
        "config": {},
        "type": "irdistance",
    },
}
