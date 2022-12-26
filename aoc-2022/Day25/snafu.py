def convert(line):
    base10 = 0
    digits = len(line)

    # Reverse string to make parsing easier
    line = line[::-1]
    for i in range(digits):
        num = 0
        if   line[i] == "-": num = -1
        elif line[i] == "=": num = -2
        else:                num = int(line[i])

        base10 += num * (5 ** i) #power operator: '**'
    
    return base10

# Implements converting from base-5 to SNAFU
# When a place is greater than 2, to covert, you have to change
#   to a negative number, and increase the previous
#   (if there is one)
# BUT, that might make the previous number greater than 2, so
#   this is called recursively until the biggest number
def updatePrev(i, base5):
    pass


# Get input
with open("Day25\input.txt", "r") as input:
    lines = input.readlines()

# Add up all the numbers
sum = 0
for line in lines:
    sum += convert(line.strip())

# Convert sum back to a SNAFU number
# First, figure out biggest place that fits (34 fits into 25's place)
# -> 34 / 5^0 > 1    34 / 5^1 > 5    34 / 5^2 > 1    34 / 5^3 < 1
#                                           ^
#num = sum
num = 198
places = 0 
while num / (5 ** places) > 1: places += 1
places -= 1 #overshot, move back one

# Convert to a base-5 intermediate
base5 = []
for place in range(places, -1, -1): #count backwards to 0
    place_amount = 5 ** place
    amt = num // place_amount
    base5.append(amt)
    num -= amt * place_amount

# 198 -> 1243
# Convert base-5 to a SNAFU number
i = 0
while i < len(base5):
    if base5[i] > 2:
        updatePrev(i, base5)
        base5[i] = base5[i] - 5 #turns 3 to -2, 4 to -1

    pass
