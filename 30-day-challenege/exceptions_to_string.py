#!/bin/python

import sys

S = raw_input().strip()

try:
  print int(S)
except Exception:
  print "Bad String"

