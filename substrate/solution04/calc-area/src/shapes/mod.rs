pub trait AreaFunc {
    fn area(&self) -> f64;
}

#[derive(Debug)]
pub struct RightTriangle {
    base: f64,
    height: f64,
}

#[derive(Debug)]
pub struct Rectangle {
    width: f64,
    height: f64,
}

#[derive(Debug)]
pub struct Circle {
    radius: f64,
}

impl RightTriangle {
    pub fn new(base: f64, height: f64) -> Self {
        RightTriangle {
            base: base,
            height: height,
        }
    }
}

impl Rectangle {
    pub fn new(width: f64, height: f64) -> Self {
        Rectangle {
            width: width,
            height: height,
        }
    }
}

impl Circle {
    pub fn new(radius: f64) -> Self {
        Circle { radius: radius }
    }
}

impl AreaFunc for RightTriangle {
    fn area(&self) -> f64 {
        (self.base * self.height) / 2.0
    }
}

impl AreaFunc for Rectangle {
    fn area(&self) -> f64 {
        self.width * self.height
    }
}

impl AreaFunc for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * self.radius.powf(2.0)
    }
}

pub fn print_area<T: AreaFunc + std::fmt::Debug>(s: &T) {
    println!("The area of shape {:?} is {}", s, s.area());
}
