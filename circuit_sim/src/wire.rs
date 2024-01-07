use crate::State;

#[derive(PartialEq, Debug)]
pub enum WireType {
    Input,
    Output,
    Connecting,
}

#[derive(PartialEq, Debug)]
pub struct Wire {
    pub wire_type: WireType,
    pub name: Option<String>,
    pub number: usize,
    pub output_gates: Option<Vec<usize>>,
    pub state: State,
}

impl Wire {
    pub fn parse(tokens: Vec<&str>) -> Wire {
        let wire_type = match tokens[0] {
            "INPUT" => WireType::Input,
            "OUTPUT" => WireType::Output,
            _ => panic!("Error in file, unknown wire type"),
        };

        let num:usize = tokens[2]
            .parse()
            .expect("The last token is the wire's number");

        Wire {
            wire_type,
            name: Some(tokens[1].to_string()),
            number: num,
            output_gates: None,
            state: State::Unknown,
        }
    }

    pub fn new_connecting(number: usize) -> Wire {
        Wire {
            wire_type: WireType::Connecting,
            name: None,
            number,
            output_gates: None,
            state: State::Unknown,
        }
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_input_wire() {
        let tokens: Vec<&str> = vec!["INPUT", "A", "3"];
        let wire = Wire::parse(tokens);

        assert_eq!(
            wire,
            Wire {
                wire_type: WireType::Input,
                name: Some("A".to_string()),
                number: 3,
                output_gates: None,
                state: State::Unknown,
            }
        );
    }

    #[test]
    fn new_output_wire() {
        let tokens: Vec<&str> = vec!["OUTPUT", "ABAB", "392"];
        let wire = Wire::parse(tokens);

        assert_eq!(
            wire,
            Wire {
                wire_type: WireType::Input,
                name: Some("ABAB".to_string()),
                number: 392,
                output_gates: None,
                state: State::Unknown,
            }
        );
    }

    #[test]
    fn new_connecting_wire() {
        assert_eq!(
            Wire::new_connecting(67242),
            Wire {
                wire_type: WireType::Connecting,
                name: None,
                number: 67242,
                output_gates: None,
                state: State::Unknown,
            }
        )
    }
}