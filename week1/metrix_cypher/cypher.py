def substract(m1 , m2):
	return [m1[0] - m2[0], m1[1] - m2[1]]


def isNotUnity(m1 , m2 ):
	unity = [1, 0, 0, 1]
	if m1[0] == unity[0] and m1[1] == unity[1] and m2[0] == unity[2] and m2[1] == unity[3]:
		return False
	return True


def sum(arr):
	sum = 0.0
	for v in arr:
		sum += v
	return sum

def readLine():
    s = input()
    s = s.split(" ")
    return [int(s[0]),int(s[1])]

def exec(top, bottom):
    result = ""
    counter = 0
    while isNotUnity(top, bottom):
        if counter >= 120:
            break
        if sum(bottom) > sum(top):
            bottom = substract(bottom, top)
            result = result + "0"
        elif sum(bottom) < sum(top):
            top = substract(top, bottom)
            result = result + "1"
        counter = counter +1
    return result

def main():
    top = readLine()
    bottom = readLine()
    result = exec(top,bottom)
    print(result)

if __name__ == "__main__":
    # execute only if run as a script
    main()