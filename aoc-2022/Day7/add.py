# Computes the total of the output for part 1. There is an easier way to do this

input = open("Day7\\nums2.txt")
lines = input.readlines()

total = 0
for line in lines:
    total += int(line.strip())

print(total)

input.close()