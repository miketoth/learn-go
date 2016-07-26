import sys

class Solution:
    # Write your code here
    def __init__(self):
        self.stack = ""
        self.queue = ""

    def pushCharacter(self, element):
        self.stack += element
 
    def enqueueCharacter(self, element):
        self.queue += element
        
    def popCharacter(self):
        temp = self.stack[-1:]
        self.stack = self.stack[0:len(self.stack)-1]
        return temp
        
    def dequeueCharacter(self):
        temp = self.queue[0]
        self.queue = self.queue[1:len(self.queue)]
        return temp

# read the string s
s=raw_input()
#Create the Solution class object
obj=Solution()   

l=len(s)
# push/enqueue all the characters of string s to stack
for i in range(l):
    obj.pushCharacter(s[i])
    obj.enqueueCharacter(s[i])

isPalindrome=True
'''
pop the top character from stack
dequeue the first character from queue
compare both the characters
''' 
for i in range(l / 2):
    if obj.popCharacter()!=obj.dequeueCharacter():
        isPalindrome=False
        break
#finally print whether string s is palindrome or not.
if isPalindrome:
    sys.stdout.write ("The word, "+s+", is a palindrome.")
else:
    sys.stdout.write ("The word, "+s+", is not a palindrome.")    
