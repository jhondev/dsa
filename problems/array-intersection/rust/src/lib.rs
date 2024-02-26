use std::collections::HashMap;

pub fn inter<'a>(left: &'a [i32], right: &'a [i32]) -> Vec<i32> {
    let lsum = summarize(left);
    let lright = summarize(right);
    let mut result = Vec::new();
    for (k, v) in lsum.iter() {
        let opt = lright.get(k);
        let num = match opt {
            Some(vr) => {
                if vr > v {
                    v
                } else {
                    vr
                }
            }
            None => continue,
        };
        for _ in 0..*num {
            result.push(*k);
        }
    }
    result
}

fn summarize(a: &[i32]) -> HashMap<i32, i32> {
    let mut map = HashMap::<i32, i32>::new();
    for &k in a {
        *map.entry(k).or_insert(0) += 1;
    }
    map
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn case1() {
        let left = [23, 3, 1, 2];
        let right = [6, 2, 4, 23];
        let result = inter(&left, &right);
        assert_eq!(result, [23, 2]);
    }

    #[test]
    fn case2() {
        let left = [1, 1, 1];
        let right = [1, 1, 1, 1];
        let result = inter(&left, &right);
        assert_eq!(result, [1, 1, 1]);
    }
}
