C, stops = [int(i) for i in input().split()]

ontrain = 0

poss = True

left, entered, stayed = [0, 0, 0]

for s in range(stops):
    left, entered, stayed = [int(i) for i in input().split()]

    ontrain -= left

    if ontrain < 0:
        poss = False
        break

    ontrain += entered
        
    if ontrain > C:
        poss = False
        break

    if stayed > 0 and ontrain < C:
        poss = False
        break

if entered != 0:
    poss = False

if stayed != 0:
    poss = False

if ontrain != 0:
    poss = False

if poss:
    print("possible")
else:
    print("impossible")