# Find the common character between the two halves of each line
# Then add up and return each character's value

input = open("Day3\input.txt", "r")
lines = input.readlines()

total = 0

for line in lines:
    found = False

    # Get first and second half
    first = line[:int(len(line)/2)]
    second = line[int(len(line)/2):]
    
    # Find common character between the halves
    for char in first:
        if char in second and not found:
            found = True
            if char.islower():              #ord gives ASCII val of char
                total += (ord(char) - 96)   #lowercase starts at 97 (a -> 1,  z -> 26)
            else:
                total += (ord(char) - 38)   #uppercase starts at 65 (A -> 27, Z -> 52)

print(total)
input.close()