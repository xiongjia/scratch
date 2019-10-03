#!/usr/bin/env python

import time
import sched
import sys
import logging
import threading
import signal
import concurrent.futures
from datetime import timedelta
from concurrent.futures import ThreadPoolExecutor
import queue


logging.basicConfig(format = '[%(asctime)s] [%(thread)d] : %(message)s')
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

def invoke():
    logger.debug("invoking")
    slow_task()

class Job(threading.Thread):
    def __init__(self, interval):
        threading.Thread.__init__(self)
        self.stopped = threading.Event()
        self.interval = interval
        self.executor = ThreadPoolExecutor(max_workers = 5)

    def stop(self):
        self.stopped.set()
        self.executor.shutdown(wait = False)
        self.join()

    def run(self):
        logger.debug("start running...")
        self.executor.submit(invoke)
        while not self.stopped.wait(self.interval.total_seconds()):
            self.executor.submit(invoke)



class ProcKilled(Exception):
    pass

def signal_handler(signum, frame):
    raise ProcKilled

q = queue.Queue()

def worker():
    logger.debug("start worker")
    while True:
        item = q.get()
        logger.debug("got work")
        if item is None:
            break
        logger.debug("work %s", item)
        invoke()
        q.task_done()

def main():
    signal.signal(signal.SIGTERM, signal_handler)
    signal.signal(signal.SIGINT, signal_handler)
    executor = ThreadPoolExecutor(5)
    executor.submit(worker)
    executor.submit(worker)
    executor.submit(worker)
    executor.submit(worker)
    executor.submit(worker)

    logger.debug("sleeping")
    try:
        while True:
            time.sleep(5)
            q.put("item")
    except:
        pass


if __name__ == "__main__":
    main()
    print("end")
