trait TrafficLightState {
    fn time(state: TrafficLight) -> u8;
}

enum TrafficLight {
    Red,
    Yellow,
    Green,
}

impl TrafficLightState for TrafficLight {
    fn time(state: TrafficLight) -> u8 {
        match state {
            TrafficLight::Red => 90,
            TrafficLight::Yellow => 3,
            TrafficLight::Green => 60,
        }
    }
}

fn main() {
    println!(
        "Traffic light red time is {} seconds.",
        TrafficLight::time(TrafficLight::Red)
    );
    println!(
        "Traffic light yellow time is {} seconds.",
        TrafficLight::time(TrafficLight::Yellow)
    );
    println!(
        "Traffic light green time is {} seconds.",
        TrafficLight::time(TrafficLight::Green)
    );
}
