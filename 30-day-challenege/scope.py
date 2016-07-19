class Difference:
    def __init__(self, a):
        self.__elements = a
    def computeDifference(self):
      current_max = 0
      for i in range(0, len(self.__elements)): 
        for j in range(1, len(self.__elements)):
          output = self.__elements[i] - self.__elements[j]
          if output < 0:
            output *= -1
          if i == 0 and j == 0:
            current_max = output
          elif output > current_max:
            current_max = output
        
      self.maximumDifference = current_max
      

# End of Difference class

_ = raw_input()
a = [int(e) for e in raw_input().split(' ')]

d = Difference(a)
d.computeDifference()

print d.maximumDifference

