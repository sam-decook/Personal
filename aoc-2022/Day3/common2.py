# Now find the common letter between each triplet of three lines
# Then add up and return the value of those letters

input = open("Day3\input.txt", "r")
lines = input.readlines()

total = 0

for i in range(0, len(lines), 3): #every third number
    found = False

    # Find common character between groups of three lines
    for char in lines[i]:
        if char in lines[i+1] and char in lines[i+2] and not found:
            found = True
            if char.islower():              #ord gives ASCII val of char
                total += (ord(char) - 96)   #lowercase starts at 97 (a -> 1,  z -> 26)
            else:
                total += (ord(char) - 38)   #uppercase starts at 65 (A -> 27, Z -> 52)

print(total)
input.close()