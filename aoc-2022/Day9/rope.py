def parseLine(line):
    dir = line[0]
    amt = line[2:].strip()
    return (dir, int(amt))

input = open("Day9\input.txt", "r")
lines = input.readlines()

head = [0,0]
tail = [0,0]
visited = {(0,0)}   #set - no duplicates allowed

for line in lines:
    dir, amt = parseLine(line)

    for x in range(amt):
        # Move head
        if   dir == "U": head[1] += 1
        elif dir == "D": head[1] -= 1
        elif dir == "L": head[0] -= 1
        elif dir == "R": head[0] += 1

        # Move tail (if it needs to)
        if abs(head[0] - tail[0]) > 1 or abs(head[1] - tail[1]) > 1:
            # Tail ends up in spot behind in the direction it moved
            # NOTE: this isn't strictly how it works, but is simpler
            if dir == "U":
                tail[0] = head[0]
                tail[1] = head[1] - 1
            elif dir == "D":
                tail[0] = head[0]
                tail[1] = head[1] + 1
            elif dir == "L":
                tail[0] = head[0] + 1
                tail[1] = head[1]
            elif dir == "R":
                tail[0] = head[0] - 1
                tail[1] = head[1]

        # Update visited - since it's a set, duplicated aren't added
        visited.add((tail[0], tail[1]))

print(str(len(visited)))
input.close()