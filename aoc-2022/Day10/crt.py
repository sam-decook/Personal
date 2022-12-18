# The CRT screen is 40 x 6 in resolution
# The clock cycle determines the pixel that is being drawn
# Now, the X register determines the position of a sprite (3 wide)
# If the pixel being drawn overlaps with the sprite, it is lit up (#)

def drawScreen(cycle, reg_x):
    # Get position in row: clock cycle starts at 1, crt position at 0
    position = (cycle - 1) % 40

    if position >= reg_x-1 and position <= reg_x+1: 
        print("#", end=" ")
    else:
        print(".", end=" ")

    if cycle % 40 == 0: print() #start drawing next row

input = open("Day10\input.txt", "r")
lines = input.readlines()

cycle = 0
reg_x = 1

for line in lines:
    # Both cases are at least one cycle
    cycle += 1
    drawScreen(cycle, reg_x)

    if line.strip() != "noop":
        cycle += 1                  #execute second cycle
        drawScreen(cycle, reg_x)    #do this first, X is updated after cycle
        a = line.split()
        reg_x += int(a[1])          #line: "addx [num]"

input.close()