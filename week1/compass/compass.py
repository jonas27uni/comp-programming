MAX = 360

start = int(input())
end = int(input())

def calcPos(s,e):
    if s<=e:
        return e-s
    return MAX - s + e

pos = calcPos(start,end)

if pos <= 180:
    print(pos)
else:
    print(pos-360)
