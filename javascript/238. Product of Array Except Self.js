/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
    const answer = [];

    let leftMul = 1;
    let rightMul = 1;
    for (let i = nums.length - 1; i >= 0; i--) {
        answer[i] = rightMul;
        rightMul *= nums[i];
    }
    for (let j = 0; j < nums.length; j++) {
        answer[j] *= leftMul;
        leftMul *= nums[j];
    }
    return answer;
};
