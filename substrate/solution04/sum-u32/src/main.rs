#[allow(dead_code)]
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

fn sum_u321(a: &[u32]) -> Option<u32> {
    let mut it = a.iter();
    it.try_fold(0u32, |acc, &x| acc.checked_add(x))
}

fn main() {
    let l = vec![1, 2, 3];
    match sum_u321(&l) {
        Some(v) => println!("sum is {}", v),
        None => println!("sum overflow"),
    }
    let l = vec![1, u32::MAX];
    match sum_u321(&l) {
        Some(v) => println!("sum is {}", v),
        None => println!("sum overflow"),
    }
}
