fn sum_u32(a: &[u32]) -> Option<u32> {
    let mut sum: u32 = 0;
    for v in a.iter() {
        match sum.checked_add(*v) {
            Some(v) => sum = v,
            None => return None,
        }
    }
    Some(sum)
}

fn main() {
    let l = vec![1, 2, 3];
    match sum_u32(&l) {
        Some(v) => println!("sum is {}", v),
        None => println!("sum overflow"),
    }
    let l = vec![1, u32::MAX];
    match sum_u32(&l) {
        Some(v) => println!("sum is {}", v),
        None => println!("sum overflow"),
    }
}
