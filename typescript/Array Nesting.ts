function arrayNesting(nums: number[]): number {
  const seen = new Set();
  let result = 0;

  nums.forEach((num) => {
    let count = 0;
    while (!seen.has(num)) {
      seen.add(num);
      count++;
      num = nums[num];
    }
    result = Math.max(result, count);
  });

  return result;
}
