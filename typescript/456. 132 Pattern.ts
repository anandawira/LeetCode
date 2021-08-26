function find132pattern(nums: number[]): boolean {
  let min = nums[0];
  let max = nums[0];

  for (let i = 1; i < nums.length; i++) {
    const cur = nums[i];
    if (cur > min && cur < max) {
      return true;
    }

    if (cur > max) {
      max = cur;
    }
  }
  return false;
}

find132pattern([1, 0, 1, -4, -3]);
