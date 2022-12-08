# Calculate score from line
#   First letter is opponent's move
#   Second letter is desired outcome
#
# opponent:  a - rock    b - paper   c - scissor
# you:       x - lose    y - draw    z - win
# lose = 0   tie = 3     win = 6

input = open("input.txt", "r")
lines = input.readlines()

score = 0

for line in lines:
    if line[0] == 'A':          #rock
        if line[2] == 'X':
            score += (3 + 0)
        elif line[2] == 'Y':
            score += (1 + 3)
        elif line[2] == 'Z':
            score += (2 + 6)
    elif line[0] == 'B':        #paper
        if line[2] == 'X':
            score += (1 + 0)
        elif line[2] == 'Y':
            score += (2 + 3)
        elif line[2] == 'Z':
            score += (3 + 6)
    elif line[0] == 'C':        #scissor
        if line[2] == 'X':
            score += (2 + 0)
        elif line[2] == 'Y':
            score += (3 + 3)
        elif line[2] == 'Z':
            score += (1 + 6)

print(score)
input.close()