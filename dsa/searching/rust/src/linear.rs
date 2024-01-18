fn search(source: &[i32], value: i32) -> Option<usize> {
    for i in 0..source.len() {
        if source[i] == value {
            return Some(i);
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn search_int() {
        let source = [1, 2, 3, 4, 5];
        assert_eq!(search(&source, 4), Some(3));
    }

    #[test]
    fn search_not_found() {
        let source = [1, 2, 3, 4, 5];
        assert_eq!(search(&source, 6), None);
    }
}
