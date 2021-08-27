use std::time::{Instant};


pub fn sol_problem001() {
    let mut sum = 0;
    let start = Instant::now();

    for i in 1..1000 { if i%3 == 0 || i%5 == 0 { sum += i; } }

    println!("Answer: {}", sum);
    println!("Time Taken: {}", start.elapsed().as_millis());
}
