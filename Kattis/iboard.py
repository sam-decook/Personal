from sys import stdin

for line in stdin:
    zeroSafe = True
    oneSafe = True

    for c in line:
        bin = ord(c)
        
        if c == '\n':
            continue


        for i in range(7):
            bit = bin & 1
            bin = bin >> 1

            if bit == 1:
                oneSafe = not oneSafe
            if bit == 0:
                zeroSafe = not zeroSafe


    if oneSafe and zeroSafe:
        print("free")
    else:
        print("trapped")
        