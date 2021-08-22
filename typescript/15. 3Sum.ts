function threeSum(nums: number[]): number[][] {
  const target = 0;
  const result = new Array();

  nums = nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length - 2; i++) {
    if (nums[i] > target) break;

    if (i > 0 && nums[i] === nums[i - 1]) continue;

    let l = i + 1;
    let r = nums.length - 1;

    while (l < r) {
      const sum = nums[i] + nums[l] + nums[r];

      if (sum === target) {
        result.push([nums[i], nums[l], nums[r]]);

        while (nums[l] === nums[l + 1]) l++;
        while (nums[r] === nums[r - 1]) r--;

        l++;
        r++;
      } else if (sum < target) {
        l++;
      } else {
        r--;
      }
    }
  }

  return result;
}
