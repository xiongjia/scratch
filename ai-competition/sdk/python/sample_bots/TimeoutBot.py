#!/usr/bin/env python

# help locate the ants.py both from batch play_one_game.bat or direct run 
# playgame.py with proper parameters
import os 
import sys
sdk_material_path = None 
while True:
    if not sdk_material_path:
        sdk_material_path = os.path.realpath('.')
    else:
        sdk_material_path = os.path.join(sdk_material_path, '..')
    ant_path = os.path.join(sdk_material_path, 'sdk\\python\\ants.py')
    if os.path.isfile(ant_path):
        break 
sys.path.append(os.path.split(ant_path)[0])

from ants import *

import time

class TimeoutBot:
    def __init__(self):
        self.gander = ['duck', 'duck', 'goose']
    def do_turn(self, ants):
        if self.gander.pop(0) == 'goose':
            sys.stderr.write("Cooking my goose...\n")
            time.sleep((ants.turntime * 2)/1000)

if __name__ == '__main__':
    try:
        import psyco
        psyco.full()
    except ImportError:
        pass
    try:
        Ants.run(TimeoutBot())
    except KeyboardInterrupt:
        print('ctrl-c, leaving ...')
