# Each line contains two ranges
# Part 1: Find the number of line in which one range envelops the other
# Part 2: Find the number of line in which the ranges overlap

def parseLine(line):
    # Line format: 53-57,37-54
    ranges = line.split(',')
    a = ranges[0].split('-')
    b = ranges[1].split('-')
    b[1] = b[1].strip()

    return (int(a[0]), int(a[1]), int(b[0]), int(b[1]))


input = open("Day4\input.txt", "r")
lines = input.readlines()

pairs_pt1 = 0
pairs_pt2 = len(lines)  # easier to substract on non-overlapping ranges

for line in lines:
    a0, a1, b0, b1 = parseLine(line)

    """------------------------------- Part 1 -------------------------------"""
    # If the start or end of the ranges are the same, one is within the other
    if a0 == b0 or a1 == b1:    pairs_pt1 += 1
    # If a starts lower and ends higher, b is within a
    elif a0 < b0 and a1 > b1:   pairs_pt1 += 1
    # If b starts lower and ends higher, a is within b
    elif b0 < a0 and b1 > a1:   pairs_pt1 += 1

    """------------------------------- Part 2 -------------------------------"""
    # If the start of one range is less than the end of the other, they overlap
    if a1 < b0 or a0 > b1:      pairs_pt2 -= 1

print(pairs_pt1)
print(pairs_pt2)
input.close()

# Note on parsing:
#   For a small number of delimiters, you can use replace
#       eg: for this, I could replace ',' with '-' if I wanted one array
#   You can also import re and call re.split, which allows you to split on multiple delimiters