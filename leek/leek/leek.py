#!/usr/bin/env python

import time
import sched
import sys
import logging

logging.basicConfig(format = '[%(msecs)d] [%(thread)d] : %(message)s')
logger = logging.getLogger()
logger.addHandler(logging.StreamHandler(sys.stdout))
logger.setLevel(logging.DEBUG)

def cost_time(func):
    def wrapper(*args, **kwargs):
        begin_ts = time.time()
        result = func(*args, **kwargs)
        end_ts = time.time()
        logger.debug('Total time taken in : %s (%d)', func.__name__, end_ts - begin_ts)
        return result
    return wrapper

@cost_time
def simple_task():
    print('simple task')
    time.sleep(3)

@cost_time
def slow_task():
    print('slow task')
    time.sleep(6)

def main():
    simple_task()
    slow_task()


if __name__ == "__main__":
    main()
