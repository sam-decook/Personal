class Monkey:
    def __init__(self, i, op, t, thr):
        self.items = i          #array of items
        self.operation = op     #operation to perform (string)
        self.test = t           #test to perform (int - all division)
        self.throw = thr        #which monkeys to throw to (tuple)
        self.num_inspected = 0  #amount of items inspected


def takeTurn(monkey, monkeys):
    while len(monkey.items) != 0:
        item = monkey.items.pop(0)

        # Increment number inspected
        monkey.num_inspected += 1

        # Perform operation on item (interest level)
        op = monkey.operation[0]
        num = monkey.operation[1]
        if (op == '+'):   item += num
        elif (op == '*'): item *= num
        elif (op == '^'): item *= item

        # Divide interest level by 3 (monkey got bored)
        item = item // 3 # '//' denotes integer division

        # Determine which monkey to throw to next
        to = monkey.throw[0] if item % monkey.test == 0 else monkey.throw[1]
        monkeys[to].items.append(item)

# Iniatialize monkeys by hand from input.txt
monkeys = [Monkey([56,52,58,96,70,75,72],    ('*', 17), 11, (2,3)),
           Monkey([75,58,86,80,55,81],       ('+',7),   3, (6,5)),
           Monkey([73,68,73,90],             ('^',2),   5, (1,7)),
           Monkey([72,89,55,51,59],          ('+',1),   7, (2,7)),
           Monkey([76,76,91],                ('*',3),  19, (0,3)),
           Monkey([88],                      ('+',4),   2, (6,4)),
           Monkey([64,63,56,50,77,55,55,86], ('+',8),  13, (4,0)),
           Monkey([79,58],                   ('+',6),  17, (1,5))]

rounds = 20    #number of rounds to simulate for part 1
for i in range(rounds):
    for monkey in monkeys:
        takeTurn(monkey, monkeys)

# The answer is the amounts of the two highest monkeys multiplied together
amounts = []
for monkey in monkeys: amounts.append(monkey.num_inspected)
amounts.sort(reverse=1)
print("Monkey business: " + str(amounts[0] * amounts[1]))