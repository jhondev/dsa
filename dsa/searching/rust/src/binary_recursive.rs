// TODO
// pub fn search(src: &[i32], value: i32) -> Option<usize> {
//     let len = src.len();
//     if len == 0 {
//         return None;
//     }
//     let mid = len / 2;
//     let idx = mid;
//     let ivalue = src[mid];
//     if value == ivalue {
//         return Some(mid);
//     }
//     let opti = if value < ivalue {
//         search(&src[0..mid], value)
//     } else {
//         search(&src[mid + 1..len], value)
//     };

//     if let Some(i) = opti {
//         return Some(idx + i + 1);
//     }

//     None
// }
