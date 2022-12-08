# Each line contains two ranges
# Find the number of line in which one range envelops the other

input = open("Day4/input.txt", "r")
lines = input.readlines()

pairs = 0

for line in lines:
    # 53-57,37-54
    ranges = line.split(',')
    a = ranges[0].split('-')
    b = ranges[1].split('-')
    b[1] = b[1].strip()

    if (int(a[0]) == int(b[0])) or (int(a[1]) == int(b[1])):
        pairs += 1
    elif (int(a[0]) < int(b[0])):
        if int(a[1]) > int(b[1]):
            pairs += 1
    elif (int(a[0]) > int(b[0])):
        if int(a[1]) < int(b[1]):
            pairs += 1

print(pairs)
input.close()

# Note on parsing:
#   For a small number of delimiters, you can use replace
#       eg: for this, I could replace ',' with '-' if I wanted one array
#   You can also import re and call re.split, which allows you to split on multiple delimiters