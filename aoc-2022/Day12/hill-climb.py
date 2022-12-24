# Find a path to the target location where the elevation change is 1 or less
#
# Solved with a breadth-first search

# Contains the information each square needs to have
class Location:
    def __init__(self, coor, height):
        self.coor = coor        # tuple - (i, j)
        self.height = height    # height (a=1 to z=26, S=1, E=26)
        self.distance = -1      # distance of path
        self.previous = None    # Previous in path
        self.target = False     # Whether this is the target location


# Updates the surrounding locations 
def search(location, maze, searchList):
    # Return location if the target is reached, None if not
    if location.target: return location

    # If moving [LRUD] is legal
    # - set current location as a previous location to next one
    # - set its distance to one more than the current
    # - add that location to the search list
    i, j = location.coor

    # Check square to the left
    if (j > 0 and maze[i][j-1].distance == -1 and 
            (maze[i][j-1].height - location.height) <= 1):
        maze[i][j-1].previous = location
        maze[i][j-1].distance = location.distance + 1
        searchList.append(maze[i][j-1])

    # Check square to the right
    if (j < len(maze[0])-1 and maze[i][j+1].distance == -1 and 
            (maze[i][j+1].height - location.height) <= 1):
        maze[i][j+1].previous = location
        maze[i][j+1].distance = location.distance + 1
        searchList.append(maze[i][j+1])

    # Check square above
    if (i > 0 and maze[i-1][j].distance == -1 and 
            (maze[i-1][j].height - location.height) <= 1):
        maze[i-1][j].previous = location
        maze[i-1][j].distance = location.distance + 1
        searchList.append(maze[i-1][j])

    # Check square below
    if (i < len(maze)-1 and maze[i+1][j].distance == -1 and 
            (maze[i+1][j].height - location.height) <= 1):
        maze[i+1][j].previous = location
        maze[i+1][j].distance = location.distance + 1
        searchList.append(maze[i+1][j])
    
    return None


# Read input from file
with open("Day12\input.txt", "r") as input:
    lines = input.readlines()

# Parse input
maze = []
for i in range(len(lines)):
    locs = []
    for j in range(len(lines[i].strip())):
        height = ord(lines[i][j]) - 96
        if lines[i][j] == 'S':
            start = (i,j)
            height = 1
        elif lines[i][j] == 'E':
            end = (i,j)
            height = 26

        locs.append( Location((i,j), height) )

    maze.append(locs)

# Correct values in start and end locations
maze[start[0]][start[1]].distance = 0
maze[end[0]][end[1]].target = True

# Find best path up hill
searchList = [maze[start[0]][start[1]]] #this stores locations to search

target = None
while target == None and len(searchList) > 0:
    # Returns none if target isn't found, Location if it is
    next = searchList.pop(0)
    target = search(next, maze, searchList)

print(str(target.distance))