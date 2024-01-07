#[derive(PartialEq, Debug)]
pub struct Circuit {
    name: String,
    input_wires: Vec<(WireType, Option(name), id)>,
    output_wire: (WireType, Option(name), id),
    gates: Vec<(GateType, Duration, Vec<usize>, usize)>,
}