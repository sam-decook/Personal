use std::{cmp::Ordering, time::Duration};

use crate::simulation::wire::State;

#[derive(PartialEq, Eq, Debug)]
pub struct Event {
    pub start: Duration,
    pub wire: usize,
    pub state: State
}

impl Ord for Event {
    fn cmp(&self, other: &Self) -> Ordering {
        other.start.cmp(&self.start)
            .then_with(|| self.wire.cmp(&other.wire))
    }
}

impl PartialOrd for Event {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}