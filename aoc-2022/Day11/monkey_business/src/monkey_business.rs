#[derive(Debug)]
struct Monkey {
    items : Vec<i64>,           //items the monkey has
    operation : (char, i32),    //the operation the monkey performs
    test : i32,                 //uses this number to test item
    throw : (i32, i32),         //determines which monkey to throw to
    num_inspected : i32,        //amount of items the monkey has inspected
}

fn build_monkey(items : Vec<i32>, operation : (char, i32), test : i32, 
        throw : (i32,i32), num_inspected : i32) {
    Monkey {
        items,
        operation,
        test,
        throw,
        num_inspected,
    }
}


fn takeTurn(monkey : Monkey, monkeys : [Monkey; 8]) {
    let mut item : i64;
    let op : char;
    let num;
    while monkey.items.len() != 0 {
        item = monkey.items.remove(0);

        // Increment number inspected
        monkey.num_inspected += 1;

        // Perform operation on item (interest level)
        op = monkey.operation[0];
        num = monkey.operation[1];
        if (op == '+')      item += num;
        else if (op == '*') item *= num;
        else if (op == '^') item *= item;

        item /= 3;

        // Determine which monkey to throw to next
        let to = if item % monkey.test == 0 {0} else {1};
        monkeys[to].items.push(item)
    }
}

fn main() {
    // Iniatialize monkeys by hand from input.txt
    let monkeys : [Monkey; 8] = [
        build_monkey(vec![56,52,58,96,70,75,72],    ('*',17), 11, (2,3)),
        build_monkey(vec![75,58,86,80,55,81],       ('+',7),   3, (6,5)),
        build_monkey(vec![73,68,73,90],             ('^',2),   5, (1,7)),
        build_monkey(vec![72,89,55,51,59],          ('+',1),   7, (2,7)),
        build_monkey(vec![76,76,91],                ('*',3),  19, (0,3)),
        build_monkey(vec![88],                      ('+',4),   2, (6,4)),
        build_monkey(vec![64,63,56,50,77,55,55,86], ('+',8),  13, (4,0)),
        build_monkey(vec![79,58],                   ('+',6),  17, (1,5))];

    let rounds = 20;  //number of rounds to simulate for part2
    for mut i in 0..range {
        for monkey in monkeys.iter() takeTurn(monkey, monkeys);
        /*
        if i % 1000 == 0 or i == 500 {
            println!("-------- After round {} --------", i + 1);
            for monkey in monkeys: println!("Business conducted: {}", monkey.num_inspected);
        }*/
    }

    // The answer is the amounts of the two highest monkeys multiplied together
    println!("Monkey business conducted:")
    for monkey in monkeys.iter() {
        println!("\t{}", monkey.num_inspected);
    }
}