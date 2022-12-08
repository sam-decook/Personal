# Each line contains two ranges
# Find the number of line in which the ranges overlap

input = open("Day4/input.txt", "r")
lines = input.readlines()

# Far easier to substract on non-overlapping ranges
pairs = len(lines)

for line in lines:
    # 53-57,37-54
    ranges = line.split(',')
    a = ranges[0].split('-')
    b = ranges[1].split('-')
    b[1] = b[1].strip()

    if (int(a[1]) < int(b[0])) or (int(a[0]) > int(b[1])):
        pairs -= 1

print(pairs)
input.close()