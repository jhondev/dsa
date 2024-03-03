pub fn sort<T: PartialOrd>(src: &mut [T]) {
    if src.is_empty() {
        return;
    }
    let pivot = src.len() / 2;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sort_1() {
        let mut src = [9, 3, 7, 4, 69, 420, 42];
        sort(&mut src);
        assert_eq!([3, 4, 7, 9, 42, 69, 420], src);
    }
}
