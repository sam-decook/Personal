use crate::State;
use std::time::Duration;

#[derive(PartialEq, Debug)]
pub enum GateType { Not, And, Or, Xor, Nand, Nor, Xnor, }

#[derive(PartialEq, Debug)]
pub struct Gate {
    pub gate_type: GateType,
    pub delay: Duration,
    pub inputs: Vec<usize>,
    pub output: usize,
    pub state: State,
    pub id: usize,
}

impl Gate {
    pub fn parse(tokens: Vec<&str>, id: usize) -> Gate {
        let gate_type = match tokens[0] {
            "NOT" => GateType::Not,
            "AND" => GateType::And,
            "OR" => GateType::Or,
            "XOR" => GateType::Xor,
            "NAND" => GateType::Nand,
            "NOR" => GateType::Nor,
            "XNOR" => GateType::Xnor,
            s => panic!("Error in file, unknown gate type: {s}"),
        };
    
        // For delay, remove "ns" and parse remaining number
        let t: u64 = tokens[1][0..tokens[1].len()-2]
            .parse()
            .expect("The delay should be in the format `[time]ns`");
    
        let delay = Duration::from_nanos(t);
    
        let mut inputs: Vec<usize> = Vec::new();
    
        for token in tokens.iter().skip(2) {
            let n: usize = token.parse()
                .expect("Gate shold have a number here");
    
            inputs.push(n);
        }
    
        let output = inputs.pop()
            .expect("There should be wire numbers");
    
        Gate {
            gate_type, 
            delay, 
            inputs, 
            output, 
            state: State::Unknown,
            id
        }
    }
}