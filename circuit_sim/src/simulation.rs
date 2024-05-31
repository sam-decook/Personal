mod wire;
mod gate;
mod event;

use wire::{Wire, State, Type};
use gate::Gate;
use event::Event;

use std::borrow::BorrowMut;
use std::cell::RefCell;
use std::collections::BinaryHeap;
use std::fs::read_to_string;
use std::rc::Rc;
use std::time::Duration;

#[derive(Debug)]
pub struct Simulation {
    wires: Vec<Rc<Wire>>,
    gates: Vec<Rc<Gate>>,
    events: BinaryHeap<Event>
}

impl Default for Simulation {
    fn default() -> Self {
        Self::new()
    }
}

impl Simulation {
    pub fn new() -> Simulation {
        let contents = read_to_string("./files/circuit2.txt").unwrap();
        let lines = contents.lines();
        
        let mut wires: Vec<Rc<Wire>> = Vec::new();
        let mut gate_lines: Vec<Vec<&str>> = Vec::new();
        
        // First pass: parse wires, skipping first line
        for line in lines.skip(1) {
            let tokens: Vec<&str> = line.split_whitespace().collect();
        
            match tokens[0] {
                "INPUT" | "OUTPUT" => wires.push(Rc::new(Wire::new(tokens))),
                _ => {
                    // Add any connecting wires
                    for token in tokens.iter().skip(2) {
                        let number: usize = token.parse().unwrap();
                        if !wires.iter().any(|wire| wire.number == number) {
                            wires.push(Rc::new(Wire {
                                wire_type: Type::Connecting,
                                number,
                                output_gates: RefCell::new(Vec::new()),
                                state: State::Unknown,
                            }));
                        }
                    }
                    
                    // Save the gate line for second pass
                    gate_lines.push(tokens);
                }
            }
        }

        wires.sort_by(|a, b| a.number.cmp(&b.number));
        
        // Second pass: parse gates and update wires with output gates
        let mut gates: Vec<Rc<Gate>> = Vec::new();
        
        for (gate_id, tokens) in gate_lines.iter().enumerate() {      
            let gate = Gate::new(tokens.to_vec(), gate_id, &wires);  
            
            for wire in &gate.as_ref().inputs {
                wire.upgrade()
                    .expect("Wire should not have been deallocated")
                    .borrow_mut()
                    .output_gates
                    .borrow_mut()
                    .push(gate.clone());
            }

            gates.push(gate);
        }
        
        // Parse vector file
        // INPUT A  6  1
        //       |  |  |-> state it is changing to
        //       |  |----> ns after start of simulation it changes
        //       |-------> the input wire
        let vector = read_to_string("./files/circuit2_v.txt").unwrap();
        let mut events: BinaryHeap<Event> = BinaryHeap::new();
        
        for line in vector.lines().skip(1) {
            let tokens: Vec<&str> = line.split_whitespace().collect();
        
            let wire = wires.iter()
                .find(|wire| {
                    match &wire.as_ref().wire_type {
                        Type::Input(name) | Type::Output(name) => {
                            name == tokens[1]
                        }
                        Type::Connecting => false
                    }
                })
                .expect("Wire not found in vector file")
                .number;
        
            let start = Duration::from_nanos(tokens[2].parse::<u64>().unwrap());
        
            let state: State = match tokens[3].parse::<u8>().unwrap() {
                0 => State::Low,
                1 => State::High,
                _ => panic!("State can only change to 1 or 0")
            };
        
            events.push(Event { start, wire, state });
        }
        
        Simulation { wires, gates, events }
    }
}
