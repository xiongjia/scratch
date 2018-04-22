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

from random import shuffle
from ants import *

import logging
logger = logging.getLogger('ConsoleLog')

class RandomBot:
    def do_turn(self, ants):
        logger.info("process guard group start...")

        destinations = []
        for a_row, a_col in ants.my_ants():
            # try all directions randomly until one is passable and not occupied
            directions = AIM.keys()
            shuffle(directions)
            for direction in directions:
                (n_row, n_col) = ants.destination(a_row, a_col, direction)
                if (not (n_row, n_col) in destinations and
                        ants.passable(n_row, n_col)):
                    ants.issue_order((a_row, a_col, direction))
                    destinations.append((n_row, n_col))
                    break
            else:
                destinations.append((a_row, a_col))

if __name__ == '__main__':
    try:
        import psyco
        psyco.full()
    except ImportError:
        pass
    try:
        Ants.run(RandomBot())
    except KeyboardInterrupt:
        print('ctrl-c, leaving ...')
