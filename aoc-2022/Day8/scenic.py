# Directional visibility function. Checks how many trees can be seen
# --> num of trees that can be seen (always at least one)
def visibleLeft(coor, lines):
    i, j = coor
    height = lines[i][j]

    num = 1
    while j - num > 0:
        if lines[i][j-num] < height: num += 1
        else: break
    
    return num

# I'm not actually sure how these are working (with num += 1 being before/after the if)
def visibleRight(coor, lines):
    i, j = coor
    height = lines[i][j]
    
    num = 1
    while j + num < len(lines[0])-1:
        if lines[i][j+num] < height: num += 1
        else: break
    
    return num

def visibleAbove(coor, lines):
    i, j = coor
    height = lines[i][j]

    num = 1
    while i - num > 0:
        if lines[i-num][j] < height: num += 1
        else: break
    
    return num

def visibleBelow(coor, lines):
    i, j = coor
    height = lines[i][j]

    num = 1
    while i + num < len(lines[0])-1:
        if lines[i+num][j] < height: num += 1
        else: break
    
    return num

# Calculates the scenic score of a tree
# This is found by multiply each of the directional scores together
def scenicScore(coor, lines):
    i, j = coor

    # Quick check if tree is on outside to speed up processing
    if i == 0 or j == 0 or i == len(lines)-1 or j == len(lines[0])-1: return 0

    return (visibleLeft(coor, lines)  * visibleRight(coor, lines) * 
            visibleAbove(coor, lines) * visibleBelow(coor, lines))

input = open("aoc-2022\Day8\input.txt", "r")
lines = input.readlines()
for i in range(len(lines)): lines[i] = lines[i].strip()

highestScore = 0
for i in range(1, len(lines)-1): #start on second and exclude last row, 
    for j in range(len(lines[0])):
        score = scenicScore((i,j), lines)
        if score > highestScore: 
            highestScore = score
            print(str(i) + "," + str(j) + " " + lines[i][j] + " " + str(score))
        
print(str(len(lines[0])))
print(highestScore)
input.close()