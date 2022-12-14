# Find the calories the elf (pt2: 3 elves) with the most calories is carrying
# Elves are separated by a blank line in the input 

input = open("input.txt", "r")
lines = input.readlines()

cals = 0
calsList = []

# Each group of calories is separated by white-space
# Add the lines together, and then append the sum to the list
for line in lines:
    if line != '\n':
        cals += int(line.strip())
    else:
        calsList.append(cals)
        cals = 0

# Sort greatest to least
calsList.sort(reverse = True)

# Answer for pt1
print(calsList[0])
# Answer for pt2
print(calsList[0] + calsList[1] + calsList[2])

input.close()
