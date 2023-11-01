//
//  -----------------------------------------
//  Primes Search - Rust version
//  -----------------------------------------
//  Author: Adam Blaszczyk
//          http://wyciekpamieci.blogspot.com
//  Date:   2023-11-01
//  -----------------------------------------
//  
//  Compilation (Windows):
//         rustc primes-search.rs -o primes-search.exe
//  Compilation (Linux):
//         rustc primes-search.rs -o primes-search
//
//  Usage:
//         primes-search
//

use std::time::SystemTime;

fn is_prime(n: u32) -> bool {
    if n < 2 {
        return false;
    } else if n == 2 {
        return true;
    } else {
        let n2 = n as f32;
        let n3 = n2.sqrt().ceil() + 1.0;
        for i in 2..n3 as u32 {
            if n % i == 0 {
                return false;
            }
        }
        return true;
    }
}

fn main() {
    let max = 8000000;
    
    println!("Checking {} numbers. Please wait...", max);
    
    let t1 = SystemTime::now();

    for i in 1..=max {
        is_prime(i);
        //println!("{}  {}", i, is_prime(i));
    }
    
    let t2 = SystemTime::now();
    let dt = t2.duration_since(t1).expect("Clock may have gone backwards.");
    println!("Done in {:?}.", dt);
}
