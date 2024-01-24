pub fn sort(src: &mut [i32]) {
    if src.is_empty() {
        return;
    }
    let mut end = src.len() - 1;
    while end > 0 {
        for i in 0..end {
            if src[i] > src[i + 1] {
                src.swap(i, i + 1)
            }
        }
        end -= 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sort_1() {
        let mut src = [1, 3, 7, 4, 2];
        sort(&mut src);
        assert_eq!([1, 2, 3, 4, 7], src);
    }

    #[test]
    fn empty() {
        let mut src = [];
        sort(&mut src);
        assert_eq!([0; 0], src);
    }
}
