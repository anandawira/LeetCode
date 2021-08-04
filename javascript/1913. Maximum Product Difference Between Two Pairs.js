/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProductDifference = function (nums) {
    let min = [1e4, 1e4];
    let max = [1, 1];

    nums.forEach((num) => {
        if (num < min[1]) {
            if (num < min[0]) {
                min = [num, min[0]];
            } else {
                min = [min[0], num];
            }
        }

        if (num > max[1]) {
            if (num > max[0]) {
                max = [num, max[0]];
            } else {
                max = [max[0], num];
            }
        }
    });

    return max[0] * max[1] - min[0] * min[1];
};
