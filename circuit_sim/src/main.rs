use circuit_sim::initialize_simulation;

fn main() {
    println!("Hello, world!");

    let sim = initialize_simulation();

    println!("\n\n\n{:#?}", sim);
}

/* Idea:
   What if I converted each of the input files to a struct
   Then I would only have to parse each file once and could start working with
   the data directly
*/

/* What if I changed the Gate and Wire structs from:
   - containing references to their linked gates and wires
   - to just containing their number
   ?

   Well, that doesn't work for gates unless I add an id field.
   I could replace the wire with their number though.
   - Pros:
     - Creating the simulation is easier, I don't have to work with complicated
       borrow semantics
     - The Gate struct wouldn't need a specified lifetime
   - Cons:
     - It's slower to find the needed wire than to have a reference
       - But maybe we could sort the wires vec in the simulation. Then, I could
         change get_wire() to check the index first

   If I keep the design (with the references), I think I will need to wrap each
   gate and wire with Rc<RefCell<>>. 
   
   Okay, that decides it, I'm moving to numbers
 */