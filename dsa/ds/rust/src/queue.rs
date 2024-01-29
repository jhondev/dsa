struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>,
}

struct Queue<T> {
    head: Option<Box<Node<T>>>,
    tail: Option<Box<Node<T>>>,
    len: usize,
}

impl<T> Queue<T> {
    fn new() -> Queue<T> {
        Queue {
            head: None,
            tail: None,
            len: 0,
        }
    }
    fn enqueue(mut self, val: &T) {
        self.head = Some(Box::new(Node::<T> {
            value: val,
            next: None,
        }));
        self.tail = Some(Box::new(Node::<T> {
            value: val,
            next: None,
        }));
        self.len += 1;
    }
    fn dequeue(mut self) {
        match self.head {
            Some(ref node) => {
                self.head = node.next;
                self.len -= 1;
            }
            _ => {}
        }
    }

    fn peek(&self) -> Option<&T> {
        match self.head {
            Some(ref node) => Some(&node.value),
            None => None,
        }
    }
}

fn insert() {
    let q = Queue::<i32>::new();
    q.enqueue(1);
}
