# Calculated score from each game
#
# opponent:  a       b       c
# you:       x       y       z
#            rock    paper   scissor
# lose = 0   tie = 3     win = 6

input = open("input.txt", "r")
lines = input.readlines()

score = 0

for line in lines:
    print(line[2])

    if line[2] == 'X':      #rock
        score += 1
        if line[0] == 'A':    #rock - tie
            score += 3
        elif line[0] == 'C':  #scissor - win
            score += 6
    elif line[2] == 'Y':    #paper
        score += 2
        if line[0] == 'B':    #paper - tie
            score += 3
        elif line[0] == 'A':  #rock - win
            score += 6
    elif line[2] == 'Z':    #scissor
        score += 3
        if line[0] == 'C':    #scissor - tie
            score += 3
        elif line[0] == 'B':  #paper - win
            score += 6

print(score)
input.close()