function getSum(a: number, b: number): number {
  return b ? getSum(a ^ b, (a & b) << 1) : a;
}
