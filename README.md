# Autonoma

The perception, planning, and control layers of the bot.

Communication to the hardware abstraction layer (HAL) is done via JSON payloads sent over TCP. This is primarly for development purposes. It makes it easier to develop when the only thing running on the bot is the HAL.

The HAL sends signal data gathered from various sensor and motor inputs to the perception layer. Then a model of the world is created or updated. The planning layer ingests the model and determines which motions to take in order to accomplish a mission. The control layer determines what the hardware (motors, sensors, etc) needs to do in order for the motions to happen.

The HAL is in it's own repo: autonoma-hal

There is also a visualization tool that can be used to see what the robot is perceiving in the autonoma-web repo.
