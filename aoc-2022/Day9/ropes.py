def parseLine(line):
    dir = line[0]
    amt = line[2:].strip()
    return (dir, int(amt))

def moveTail(head, tail, dir):
    if abs(head[0] - tail[0]) > 1 or abs(head[1] - tail[1]) > 1:
        # The trailing knot moves to be one spot behind the one ahead
        # If on same row/column, it simply moves up/down/left/right
        x_diff = 0
        y_diff = 0
        
        # If not, then it moves diagonally
        if head[0] < tail[0]: x_diff = -1
        if head[0] > tail[0]: x_diff = 1
        if head[1] < tail[1]: y_diff = -1
        if head[1] > tail[1]: y_diff = 1

        tail[0] += x_diff
        tail[1] += y_diff
    
    return tail

input = open("Day9\input.txt", "r")
lines = input.readlines()

knots = [[0,0], [0,0], [0,0], [0,0], [0,0], [0,0], [0,0], [0,0], [0,0], [0,0]]
visited = {(0,0)}   #set - no duplicates allowed


for line in lines:
    dir, amt = parseLine(line)

    for x in range(amt):
        # Move head
        if   dir == "U": knots[0][1] += 1
        elif dir == "D": knots[0][1] -= 1
        elif dir == "L": knots[0][0] -= 1
        elif dir == "R": knots[0][0] += 1

        # Move tails (if it needs to)
        for x in range(1, len(knots)):
            knots[x] = moveTail(knots[x - 1], knots[x], dir)

        # Update visited - since it's a set, duplicated aren't added
        visited.add((knots[9][0], knots[9][1]))

print(str(len(visited)))
input.close()