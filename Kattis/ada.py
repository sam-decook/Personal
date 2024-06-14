def allZeros(diffs):
    for n in diffs:
        if n != 0:
            return False
    return True

def getDifferences(differences):
    last = len(differences) - 1
    diffLen = len(differences[last]) - 1
    diffs = [differences[last][i+1] - differences[last][i] for i in range(diffLen)]

    if allZeros(diffs):
        return
    else:
        differences.append(diffs)
        getDifferences(differences)

def getNext(differences):
    nextNum = 0
    for line in differences:
        nextNum += line[-1]
    return nextNum

seq = [int(n) for n in input().split()]

differences = []
differences.append(seq[1:])

getDifferences(differences)

degree = len(differences) - 1

print(f"{degree} {getNext(differences)}")