function twoSum(nums: number[], target: number): number[] {
  const previousNumbers = new Map();

  for (let i = 0; i < nums.length; i++) {
    const curNum = nums[i];
    if (previousNumbers.has(target - curNum)) {
      return [i, previousNumbers.get(target - curNum)];
    } else {
      previousNumbers.set(curNum, i);
    }
  }

  return [];
}
