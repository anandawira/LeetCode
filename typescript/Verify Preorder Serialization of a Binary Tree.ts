function isValidSerialization(preorder: string): boolean {
  const preorderList = preorder.split(',');

  let slot = 1;
  for (const node of preorderList) {
    if (slot === 0) return false;

    if (node === '#') {
      slot -= 1;
    } else {
      slot += 1;
    }
  }

  return slot === 0;
}
