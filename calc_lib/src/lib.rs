

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(1, 2), 3);
    }

    #[test]
    fn test_subtract() {
        assert_eq!(subtract(2, 1), 1);
    }

    #[test]
    fn test_multiply() {
        assert_eq!(multiply(2, 3), 6);
    }

    #[test]
    fn test_divide() {
        assert_eq!(divide(6, 3), 2);
    }
}

#[no_mangle]
pub extern "C" fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[no_mangle]
pub extern "C" fn subtract(a: i32, b: i32) -> i32 {
    a - b
}

#[no_mangle]
pub extern "C" fn multiply(a: i32, b: i32) -> i32 {
    a * b
}

#[no_mangle]
pub extern "C" fn divide(a: i32, b: i32) -> i32 {
    a / b
}

// #[no_mangle]
// https://rs.nkmk.me/rust-division-by-zero/
// pub extern "C" fn divide(a: i32, b: i32) -> Result<i32, &'static str> {
//     if b == 0 {
//         Err("Cannot divide by zero")
//     } else {
//         Ok(a / b)
//     }
// }
