#!/usr/bin/env python

from mockito import mock, verify

import unittest

from hello import hello


class HelloWorldTest(unittest.TestCase):
  def test_should_issue_hello_world_message(self):
    out = mock()
    hello(out)
    verify(out).write("hello pyb\n")

