/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {

  for ( const [idx, val] of nums.entries()) {
    const idx2 = nums.indexOf(target - val, idx + 1);

    if (idx2 != -1) {
      return [idx, idx2];
    }
  }
};