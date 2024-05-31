use std::cell::RefCell;
use std::rc::Rc;
use std::fmt;

use super::gate::Gate;

#[derive(PartialEq, Debug)]
pub enum Type {
    Input(String),
    Output(String),
    Connecting
}

#[derive(PartialEq, Eq, Debug)]
pub enum State { Low, High, Unknown }

pub struct Wire {
    pub wire_type: Type,
    pub number: usize,
    pub output_gates: RefCell<Vec<Rc<Gate>>>,
    pub state: State,
}

impl Wire {
    pub fn new(tokens: Vec<&str>) -> Wire {
        let wire_type = match tokens[0] {
            "INPUT" => Type::Input(tokens[1].to_string()),
            "OUTPUT" => Type::Output(tokens[1].to_string()),
            _ => panic!("Error in file, unknown wire type"),
        };

        let num:usize = tokens[2]
            .parse()
            .expect("The last token is the wire's number");

        Wire {
            wire_type,
            number: num,
            output_gates: RefCell::new(Vec::new()),
            state: State::Unknown,
        }
    }
}

impl fmt::Debug for Wire {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Wire")
         .field("number", &self.number)
         .field("output_gates", &self.output_gates)
         .finish()
    }
}