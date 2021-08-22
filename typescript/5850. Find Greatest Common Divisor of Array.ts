function findGCD(nums: number[]): number {
  const euclidean = (a: number, b: number): number => {
    if (b === 0) return a;

    return euclidean(b, a % b);
  };

  let min = 1000;
  let max = 1;

  nums.forEach((num: number) => {
    min = Math.min(min, num);
    max = Math.max(max, num);
  });
  return euclidean(min, max);
}
