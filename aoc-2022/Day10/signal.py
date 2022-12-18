input = open("Day10\input.txt", "r")
lines = input.readlines()

cycle = 0
reg_x = 1

# Add up the value of the register multiplied by the cycle
strengths = 0

for line in lines:
    cycle += 1
    if cycle % 40 == 20: strengths += cycle * reg_x

    if line.strip() != "noop":
        cycle += 1
        if cycle % 40 == 20: strengths += cycle * reg_x

        a = line.split()
        reg_x += int(a[1]) #line: "addx [num]"

print(str(strengths))
input.close()