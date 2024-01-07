mod wire;
mod gate;

use crate::wire::Wire;
use crate::gate::Gate;

use std::fs::read_to_string;

// The wires and gates will be stored in vectors
// Wires will be access by their number. Since the input can skip some
//   numbers (ie. 1, 2, 4, 6), we use Option<Wire>
#[derive(Debug)]
pub struct Simulation {
    wires: Vec<Wire>,
    gates: Vec<Gate>,
}

impl Default for Simulation {
    fn default() -> Self {
        Self::new()
    }
}

impl Simulation {
    pub fn new() -> Simulation {
        Simulation {
            wires: Vec::new(),
            gates: Vec::new(),
        }
    }

    pub fn add_wires(&mut self, wires: Vec<Wire>) {
        for wire in wires {
            self.wires.push(wire);
        }
    }

    pub fn get_wire(&self, num: usize) -> Result<&Wire, &'static str> {
        self.wires.iter().find(|wire| wire.number == num)
        .ok_or("No wire found")
    }

    pub fn add_gates(&mut self, gates: Vec<Gate>) {
        for gate in gates {
            self.gates.push(gate);
        }
    }
}

#[derive(PartialEq, Debug)]
pub enum State {
    Low,
    High,
    Unknown,
}

/* The input and/or output wires for a gate are not necessarily described in 
   the input/output wire section. To make parsing easier, a two-pass approach
   will be used:
   - pass 1: parse the input, output, and any internal wires in the gate lines
   - pass 2: parse the gates
*/
pub fn initialize_simulation() -> Simulation {
    let mut sim = Simulation::new();

    let contents = read_to_string("./files/circuit2.txt").unwrap();
    let lines = contents.lines();

    let mut wires: Vec<Wire> = Vec::new();
    let mut gate_lines: Vec<&str> = Vec::new();

    // First pass: parse wires, skipping first line
    for line in lines.skip(1) {
        let tokens: Vec<&str> = line.split_whitespace().collect();

        match tokens[0] {
            "INPUT" | "OUTPUT" => wires.push(Wire::parse(tokens)),
            // Gate: save the line for later and add any connecting wires
            _ => {
                gate_lines.push(line);
                
                for token in tokens.iter().skip(2) {
                    let num: usize = token.parse().unwrap();
                    if !wires.iter().any(|wire| wire.number == num) {
                        wires.push(Wire::new_connecting(num));
                    }
                }
            }
        }
    }

    sim.add_wires(wires);

    // Second pass: parse gates
    let mut gates: Vec<Gate> = Vec::new();
    
    for (id, line) in gate_lines.iter().enumerate() {
        let tokens: Vec<&str> = line.split_whitespace().collect();

        gates.push(Gate::parse(tokens, id));
    }

    sim.add_gates(gates);

    // Finish the wires by adding the gates they output to
    // TODO: I'm copying a lot of code, see if we can fold this into the previous for loop
    for (id, line) in gate_lines.iter().enumerate() {
        let tokens: Vec<&str> = line.split_whitespace().collect();

        for token in tokens.iter().skip(2).take(tokens.len()-1) {
            let wire_num: usize = token.parse().unwrap();
        }
    }

    sim
}