use std::time::Duration;
use std::rc::{Rc, Weak};
use std::fmt;

use super::wire::Wire;

#[derive(PartialEq, Debug)]
pub enum Type { Not, And, Or, Xor, Nand, Nor, Xnor }

pub struct Gate {
    pub gate_type: Type,
    pub delay: Duration,
    pub inputs: Vec<Weak<Wire>>,
    pub output: Rc<Wire>,
    pub id: usize,
}

impl Gate {
    pub fn new(tokens: Vec<&str>, id: usize, wires: &[Rc<Wire>]) -> Rc<Gate> {
        let gate_type = match tokens[0] {
            "NOT" => Type::Not,
            "AND" => Type::And,
            "OR" => Type::Or,
            "XOR" => Type::Xor,
            "NAND" => Type::Nand,
            "NOR" => Type::Nor,
            "XNOR" => Type::Xnor,
            s => panic!("Error in file, unknown gate type: {s}"),
        };
    
        // For delay, remove "ns" and parse remaining number
        let t: u64 = tokens[1][0..tokens[1].len()-2]
            .parse()
            .expect("The delay should be in the format `[time]ns`");
    
        let delay = Duration::from_nanos(t);
    
        // Collect all wires, output wire is the last one
        let mut inputs: Vec<Weak<Wire>> = Vec::new();
    
        for token in tokens.iter().skip(2) {
            let n: usize = token.parse()
                .expect("Gate should have a number here");
    
            let wire = wires.iter()
                .find(|wire| wire.number == n)
                .expect("Wire should exist")
                .clone();

            inputs.push(Rc::downgrade(&wire));
        }
    
        let output = inputs.pop()
            .expect("There should be wire numbers")
            .upgrade()
            .expect("The wire should not have been deallocated");
    
        Rc::new(Gate {
            gate_type, 
            delay, 
            inputs, 
            output, 
            id
        })
    }
}

impl fmt::Debug for Gate {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Gate")
         .field("type", &self.gate_type)
         .field("delay", &self.delay)
         .field("output", &self.output)
         .finish()
    }
}