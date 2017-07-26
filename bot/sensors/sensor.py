#!/usr/bin/evn python3


class Sensor(object):

    def __init__(self, metadata, config):
        initial_metadata = {
            'id': "no-id",
            'xOffset': 0.0,
            'yOffset': 0.0,
            'maxDistance': 0.0,
            'minDistance': 0.0,
            'coneWidth': 0.0,
            'inclusive': False,
            'angleOffset': 0.0,
        }
        self.metadata = initial_metadata
        self.metadata.update(metadata)
        self.config = config

    def read(self):
        return 0.0
