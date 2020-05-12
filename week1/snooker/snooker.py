import numpy as np

RED = 1
YELLOW = 2
GREEN = 3
BROWN = 4
BLUE = 5
PINK = 6
BLACK = 7

def getNumber(i):
    if i =="red":
        return RED
    elif i=="yellow":
        return YELLOW
    elif i=="green":
        return GREEN
    elif i=="brown":
        return BROWN
    elif i=="blue":
        return BLUE
    elif i=="pink":
        return PINK
    else:
        return BLACK

balls = int(input())
s = np.zeros(balls, dtype=np.int)

for i in range(balls):
    s[i] = getNumber(input())
    
occurrences = np.count_nonzero(s == 1)
max = np.max(s)
score = (max+1)* occurrences
s[s==1] = 0
print(score+np.sum(s))