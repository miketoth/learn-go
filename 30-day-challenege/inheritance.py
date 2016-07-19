class Person: #wish i could change this
    def __init__(self, firstName, lastName, idNumber):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNumber
    def printPerson(self):
        print "Name:", self.lastName + ",", self.firstName
        print "ID:", self.idNumber


class Student(Person):
    def __init__(self, firstName, lastName, idNum, scores):
      #super(Student, self).__init__(firstName, lastName, idNum)
      # WHY CANT I USE SUPER?
      self.firstName = firstName
      self.lastName = lastName
      self.idNumber = idNum
      self.scores = scores
    def calculate(self):
      total = 0
      count = 0
      for i in scores:
        total += i
        count += 1
      total = total / count

      if total >= 90:
        return 'O'
      elif total >= 80:
        return 'E'
      elif total >= 70:
        return 'A'
      elif total >= 55:
        return 'P'
      elif total >= 40:
        return 'D'
      else:
        return 'T'

line = raw_input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(raw_input()) # not needed for Python
scores = map(int, raw_input().split())
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print "Grade:", s.calculate()
