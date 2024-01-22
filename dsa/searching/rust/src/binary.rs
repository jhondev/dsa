pub fn search(src: &[i32], value: i32) -> Option<usize> {
    let len = src.len();
    if len == 0 {
        return None;
    }
    let mut start = 0;
    let mut end = len;
    while start < end {
        let mid = start + (end - start) / 2;
        let ivalue = src[mid];
        if ivalue == value {
            return Some(mid);
        }
        if ivalue < value {
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
