mod shapes;

use shapes::{print_area, Circle, Rectangle, RightTriangle};

fn main() {
    let s1 = RightTriangle::new(3.0, 4.0);
    let s2 = Rectangle::new(3.0, 4.0);
    let s3 = Circle::new(5.0);
    print_area(&s1);
    print_area(&s2);
    print_area(&s3);
}
