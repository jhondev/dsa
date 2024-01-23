pub fn search<T: PartialOrd + Copy>(src: &[T], value: T) -> Option<usize> {
    let mut start = 0;
    let mut end = src.len();
    while start < end {
        let mid = start + (end - start) / 2;
        if src[mid] == value {
            return Some(mid);
        }
        if src[mid] < value {
            start = mid + 1;
        } else {
            end = mid;
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn search_empty_src() {
        assert_eq!(search(&[], 1), None);
    }

    #[test]
    fn search_int_right() {
        let src = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
        let idx = search(&src, 11).unwrap();
        assert_eq!(idx, 10);
    }

    #[test]
    fn search_int_left() {
        let src = [1, 2, 3, 4, 5, 6, 7];
        let idx = search(&src, 1).unwrap();
        assert_eq!(idx, 0);
    }

    #[test]
    fn search_not_found_right() {
        let src = [1, 2, 3];
        let idx = search(&src, 4);
        assert_eq!(idx, None);
    }

    #[test]
    fn search_not_found_left() {
        let src = [1, 2, 3];
        let idx = search(&src, 0);
        assert_eq!(idx, None);
    }
}
