function reverse(head) {
  let prev;
  let current = head;
  let nextTemp;
  while (current !== null) {
    nextTemp = current.next;
    current.next = prev;
    prev = current;
    current = nextTemp;
  }
  return prev;
}

export default reverse;