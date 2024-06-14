N, M = [int(i) for i in input().split()]

map = [{} for i in range(N+1)]


for i in range(M):
    s1, s2  = [int(i) for i in input().split()]

    map[s1][s2] = i+1
    map[s2][s1] = i+1

valid = True

k = list(map[1].keys())

for s in range(1, N+1):
    target = s+1 if s != N else 1

    if target not in list(map[s].keys()):
        valid = False
        break
if valid:
    for s in range(1, N+1):
        target = s+1 if s != N else 1

        print(map[s][target])
else:
    print("impossible")

