#!/bin/python

import sys

class Calculator(object):

  def power(self,n,p):
    output = 1
    if n < 0 or p < 0:
      raise ValueError("n and p should be non-negative")
    while p > 0:
      output *= n
      p -= 1
    return output
    
myCalculator=Calculator()
T=int(raw_input())
for i in range(T):
    n,p = map(int, raw_input().split())
    try:
        ans=myCalculator.power(n,p)
        print ans
    except Exception,e:
        print e    
    

