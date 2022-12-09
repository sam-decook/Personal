# Directional visibility function. Checks every tree in that direction
# --> True: all are shorter, visible
# --> False: one or more are taller, not visible
def visibleLeft(coor, lines):
    i, j = coor
    height = lines[i][j]
    while j > 0:
        # If neighbor is taller, tree isn't visible in that direction
        if lines[i][j-1] >= height: return False
        j -= 1
    
    # If all trees to the left are shorter, this tree is visible
    return True

def visibleRight(coor, lines):
    i, j = coor
    height = lines[i][j]
    while j < len(lines[0])-1:
        if lines[i][j+1] >= height: return False
        j += 1
    
    return True

def visibleAbove(coor, lines):
    i, j = coor
    height = lines[i][j]
    while i > 0:
        if lines[i-1][j] >= height: return False
        i -= 1
    
    return True

def visibleBelow(coor, lines):
    i, j = coor
    height = lines[i][j]
    while i < len(lines[0])-1:
        if lines[i+1][j] >= height: return False
        i += 1
    
    return True

# General visibility function. Check if the tree is visible
# --> True: visible from one or more directions
# --> False: not visible from any direction
# MAJOR NOTE: coordinates are described with (i,j)
#             i is the row in descending order (0 is top)
#             j is the column, same as x
#             It may be better not to refer to it as a coordinate, as it's not (x,y)
def visible(coor, lines):
    # Quick check if tree is on rim to speed up processing
    i, j = coor

    if i == 0 or j == 0 or i == len(lines)-1 or j == len(lines[0])-1: return True

    return (visibleLeft(coor, lines)  or visibleRight(coor, lines) or 
            visibleAbove(coor, lines) or visibleBelow(coor, lines))

input = open("aoc-2022\Day8\input.txt", "r")
lines = input.readlines()
for i in range(len(lines)): lines[i] = lines[i].strip()

numVisible = 0
for i in range(len(lines)):
    for j in range(len(lines[0])):
        if visible((i,j), lines): numVisible += 1
        
print(numVisible)
input.close()