while True:
    numHeads, numKnights = [int(x) for x in input().split()]

    if numHeads == 0 and numKnights == 0:
        break

    heads = []
    for i in range(numHeads):
        heads.append(int(input()))

    knights = []
    for i in range(numKnights):
        knights.append(int(input()))

    if numKnights < numHeads:
        print("Loowater is doomed!")
        break
    
    doomed = False

    heads.sort()
    knights.sort()

    cost = 0
    while len(heads) > 0 and len(knights) > 0 and not doomed:
        head = heads.pop(0)
        knight = knights.pop(0)

        while knight < head and len(knights) > 0:
            knight = knights.pop(0)

        if len(knights) < len(heads):
            doomed = True
        else:
            cost += knight
    
    if doomed:
        print("Loowater is doomed!")
    else:
        print(cost)

"""
2 2
4
5
4
5
2 3
5
6
7
8
4
2 1
5
5
10
3 3
5
11
8
6
9
7
0 0
"""