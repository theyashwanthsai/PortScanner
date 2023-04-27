use std::net::TcpStream;
use std::time::Duration;

fn scan_port(host: &str, port: u16, timeout: u64) -> bool {
    let addr = format!("{}:{}", host, port);
    match TcpStream::connect_timeout(&addr, Duration::from_secs(timeout)) {
        Ok(_) => true,
        Err(_) => false,
    }
}

fn main() {
    let host = "example.com";
    let timeout = 3; // in seconds
    for port in 1..65535 {
        if scan_port(host, port, timeout) {
            println!("Port {} is open", port);
        }
    }
}
