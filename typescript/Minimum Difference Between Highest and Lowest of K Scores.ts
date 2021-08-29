function minimumDifference(nums: number[], k: number): number {
  if (k === 1) return 0;
  nums.sort((a, b) => a - b);

  let min = Infinity;
  for (let i = 0; i <= nums.length - k; i++) {
    min = Math.min(min, nums[i + k - 1] - nums[i]);
  }

  return min;
}

// minimumDifference([9, 4, 1, 7], 2);
