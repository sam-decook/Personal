def floyd_warhsall(weight):
    n = len(weight)
    for k in range(n):
        for i in range(n):
            for j in range(n):
                inter = weight[i][k] + weight[k][j]
                if inter < weight[i][j]:
                    weight[i][j] = inter

n, k = [int(x) for x in input().split()]

weight = [[float('inf')] * n for _ in range(n)]

for i in range(n):
    weight[i][i] = 0

for i in range(n-1):
    n1, n2, cost = [int(x) for x in input().split()]

    weight[n2][n1] = cost
    weight[n1][n2] = cost

floyd_warhsall(weight)

leaves = [int(x) for x in input().split()]

cost = 0
curr = 0
for leaf in leaves:
    cost += weight[curr][leaf]
    curr = leaf

cost += weight[curr][0]

print(cost)
