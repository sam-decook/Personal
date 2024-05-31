mod simulation;

use simulation::Simulation;

fn main() {
    let sim = Simulation::new();

    println!("{:#?}", sim);
}