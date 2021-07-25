/**
 * @param {number[]} nums
 * @return {number}
 */
var dominantIndex = function (nums) {
    const maxNum = Math.max(...nums);
    const maxNumIdx = nums.indexOf(maxNum);

    nums.splice(maxNumIdx, 1);

    if (maxNum >= Math.max(...nums) * 2) {
        return maxNumIdx;
    } else {
        return -1;
    }
};
