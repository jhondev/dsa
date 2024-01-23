#[allow(clippy::manual_find)]
pub fn search<T: PartialOrd + Copy>(source: &[T], value: T) -> Option<usize> {
    for (i, &num) in source.iter().enumerate() {
        if num == value {
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
