# Creates a tree structure of the directories, omitting files
class Directory():
    def __init__(self, dir, parentDir):
        self.parent = parentDir
        self.name = dir
        self.subDirectories = []

        # On first pass, this is the size of its immediate files
        # On second pass, this is corrected to be the size of all files it contains
        self.size = 0


# Corrects the size of all directories by adding the size of subdirectories
def directorySize(dir):
    # Add the subdirectories' sizes to its size
    for subDir in dir.subDirectories:
        dir.size += directorySize(subDir)

    # Return updated size for:
    #   parent directory if not root
    #   total fs size if root
    return dir.size

# Prints directory structure and size recursively
# Originally mostly for debugging, then used to find all
#   directories of a certain size for parts 1 and 2
def printDirSize(dir, indent):
    # Prints out all directories and their sizes
    #print((" " * indent) + dir.name + ": " + str(dir.size))

    # Part 1: Prints out all directories smaller than 100,000
    #if dir.size < 100000:
    # Part 2: Prints out all directories larger than 6,552,309
    if dir.size > 6552309:
        print(str(dir.size))

    for d in dir.subDirectories:
        printDirSize(d, indent + 4)



input = open("Day7\\input.txt", "r")
lines = input.readlines()

# Build directory structure
rootDir = None
currDir = None

for line in lines:
    tokens = line.split()
    tokens[-1].strip()

    if tokens[0] == '$' and tokens[1] == "cd":      # $ cd [dir]
        if tokens[2] == "/":
            rootDir = Directory("/", None)
            currDir = rootDir
        elif tokens[2] == "..":
            currDir = currDir.parent
        else:
            parent = currDir
            currDir = Directory(tokens[2], parent)
            parent.subDirectories.append(currDir)

    elif tokens[0] != "dir" and tokens[0] != "$":   # [size] [filename]
        currDir.size += int(tokens[0])
    
    # There is no need to do anything for "$ ls" or directory listings

# Recurse through filesystem adding size of subdirectories to parent directories
directorySize(rootDir)

# Part 1: modify printDirSize to print directories < 100,000
# Part 2: modify printDirSize to print directories > 6,552,309, then find smallest
#   Total disk space: 70,000,000
#   We have used: 46,552,309
#   We need total unused space to be > 30,000,000
#   Find the smallest directory we can delete to reach that
#   We need to delete at least 6,552,309

# Print out updated (accurate) directory sizes)
printDirSize(rootDir, 0)

input.close()