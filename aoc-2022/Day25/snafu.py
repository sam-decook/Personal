"""------------------------------- Functions -------------------------------"""
# Converts a SNAFU number (modified base-5) to base-10 (decimal)
def convert(line):
    base10 = 0
    digits = len(line)

    # Reverse string to make parsing easier - index is power (5 ^ i)
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
# If the first (biggest) number is greater than 2, a new place is added
#   Then, this will return 1, otherwise, it returns 0
def update(i):
    # The only clean way to modify the array in recursion
    global base5    

    # If 0,1,2, no changes need to be made
    if (base5[i]) < 3: return 0

    # Keep track of this for the original loop, otherwise the index
    #   will become out-of-date and point to the wrong number
    added = 0

    # If there is no previous number, we need to add one
    if i == 0: 
        base5.insert(i, 0)
        i += 1
        added = 1
    
    base5[i] = base5[i] - 5 # Turn 3 to -2, 4 to -1
    base5[i-1] += 1         # Increment previous

    if base5[i-1] < 3: return added         # Done with recursion
    else:              return update(i-1)   # Need to continue fixing


"""---------------------------------- Code ----------------------------------"""
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
num = sum
places = 0 
while num / (5 ** places) >= 1: places += 1
places -= 1 #overshot, move back one

# Convert to a base-5 intermediate
base5 = []
for place in range(places, -1, -1): #count backwards to 0
    place_amount = 5 ** place
    amt = num // place_amount
    base5.append(amt)
    num -= amt * place_amount

# Convert base-5 to a SNAFU number
i = 0
while i < len(base5):
    placeAdded = update(i)
    i += 1 + placeAdded

# Convert to string and change -1 and -2 to '-' and '='
snafu = ""
for num in base5:
    if num == -1:   snafu += '-'
    elif num == -2: snafu += '='
    else:           snafu += str(num)

print(snafu)