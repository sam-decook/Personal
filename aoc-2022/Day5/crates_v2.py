input = open("Day5\input.txt", "r")
lines = input.readlines()

'''
Initial input:
                    [L]     [H] [W]
                [J] [Z] [J] [Q] [Q]
[S]             [M] [C] [T] [F] [B]
[P]     [H]     [B] [D] [G] [B] [P]
[W]     [L] [D] [D] [J] [W] [T] [C]
[N] [T] [R] [T] [T] [T] [M] [M] [G]
[J] [S] [Q] [S] [Z] [W] [P] [G] [D]
[Z] [G] [V] [V] [Q] [M] [L] [N] [R]
 1   2   3   4   5   6   7   8   9 
'''

# First is blank to support one-indexing
supplies = [[],
            ["Z", "J", "N", "W", "P", "S"],
            ["G", "S", "T",],
            ["V", "Q", "R", "L", "H"],
            ["V", "S", "T", "D"],
            ["Q", "Z", "T", "D", "B", "M", "J"],
            ["M", "W", "T", "J", "D", "C", "Z", "L"],
            ["L", "P", "M", "W", "G", "T", "J"],
            ["N", "G", "M", "T", "B", "F", "Q", "H"],
            ["R", "D", "G", "C", "P", "B", "Q", "W"]]

for line in lines:
    #  move [num] from [src] to [dst]
    tokens = line.split()
    num = int(tokens[1])
    src = int(tokens[3])
    dst = int(tokens[5])

    for x in range(num):
        tmp = supplies[src].pop(-num+x)
        supplies[dst].append(tmp)

for x in range(1, 10):
    print(supplies[x][-1], end='')

input.close()