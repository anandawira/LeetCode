class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for idx, val in enumerate(nums):
            try:
                idx2 = nums.index(target-val, idx+1)
                
                return [idx, idx2]
            except Exception:
                pass