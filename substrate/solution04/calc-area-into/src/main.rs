use std::f64::consts::PI;

#[derive(Debug)]
pub struct Circle<T> {
    pub r: T,
}

#[derive(Debug)]
pub struct Triangle<T, U, V> {
    pub a: T,
    pub b: U,
    pub c: V,
}

#[derive(Debug)]
pub struct Square<T> {
    pub a: T,
}

pub trait AreaCalculate {
    fn area(&self) -> f64;
}

impl<T> AreaCalculate for Circle<T>
where
    T: Copy + Into<f64>,
{
    fn area(&self) -> f64 {
        let r: f64 = self.r.into();
        PI * r * r
    }
}

impl<T> AreaCalculate for Square<T>
where
    T: Copy + Into<f64>,
{
    fn area(&self) -> f64 {
        let a: f64 = self.a.into();
        a * a
    }
}

impl<T, U, V> AreaCalculate for Triangle<T, U, V>
where
    T: Copy + Into<f64>,
    U: Copy + Into<f64>,
    V: Copy + Into<f64>,
{
    fn area(&self) -> f64 {
        let a: f64 = self.a.into();
        let b: f64 = self.b.into();
        let c: f64 = self.c.into();
        let p = (a + b + c) / 2.0;

        (p * (p - a) * (p - b) * (p - c)).sqrt()
    }
}

pub fn area<T: AreaCalculate>(shape: &T) -> f64 {
    shape.area()
}

fn main() {
    let circle = Circle { r: 10u8 };
    println!("{:#?} \n area {}\n", circle, area(&circle));

    let triangle = Triangle {
        a: 3u32,
        b: 4.0f32,
        c: 5.00000000000001f64,
    };
    println!("{:#?} \n area {}\n", triangle, area(&triangle));

    let square = Square { a: 10u32 };
    println!("{:#?} \n area {}\n", square, area(&square));
}
