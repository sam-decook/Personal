# Each line contains two ranges
# Find the number of line in which the ranges overlap

def parseLine(line):
    # 53-57,37-54
    ranges = line.split(',')
    a = ranges[0].split('-')
    b = ranges[1].split('-')
    b[1] = b[1].strip()

    return (int(a[0]), int(a[1]), int(b[0]), int(b[1]))


input = open("Day4/input.txt", "r")
lines = input.readlines()

# Far easier to substract on non-overlapping ranges
pairs = len(lines)

for line in lines:
    a0, a1, b0, b1 = parseLine(line)

    # If the start of one range is less than the end of the other, they overlap
    if a1 < b0 or a0 > b1: pairs -= 1

print(pairs)
input.close()