#!/usr/bin/evn python3


class Sensor(object):

    def __init__(self, metadata, config):
        """ metadata =
        {
            'id': id,
            'xOffset': x_offset,
            'yOffset': y_offset,
            'maxDistance': max_dist,
            'minDistance': min_dist,
            'coneWidth': cone_width,
            'inclusive': inclusive,
            'angleOffset': angle_offset,
        }
        """
        self.metadata = metadata
        self.config = config

    def read(self):
        return 0.0
