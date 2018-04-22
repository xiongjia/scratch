#!/usr/bin/env python

import os, logging

def initLogging():
    logLevel = logging.DEBUG
    log_path = os.path.join(os.getcwd(), 'scratch_ant.log')
    try:
        os.remove(log_path)
    except FileNotFoundError:
        pass

    logging.basicConfig(level = logLevel,
        format = '%(asctime)s ' + 
                 '%(funcName)-10s[%(lineno)-3d] %(levelname)-5s: %(message)s ',
        datefmt = '%H:%M:%S',
        filename = log_path,
        filemode = 'a')

    logger = logging.getLogger('main')
    logger.setLevel(logLevel)

    handler = logging.StreamHandler()
    handler.setLevel(logging.DEBUG)
    logger.addHandler(handler)

