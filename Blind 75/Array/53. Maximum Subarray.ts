function maxSubArray(nums: number[]): number {
  let subarraySum = [...nums];

  for (let i = 1; i < subarraySum.length; i++) {
    subarraySum[i] = Math.max(subarraySum[i], subarraySum[i] + subarraySum[i - 1]);
  }
  return Math.max(...subarraySum)
}
