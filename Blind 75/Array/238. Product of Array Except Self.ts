// function productExceptSelf(nums: number[]): number[] {
//   const answer = new Array(nums.length);

//   let leftMul = 1;
//   let rightMul = 1;

//   // product of elements on right side of index
//   for (let i = nums.length - 1; i >= 0; i--) {
//     answer[i] = rightMul;
//     rightMul *= nums[i];
//   }

//   // product of elements on left side of index
//   for (let j = 0; j < nums.length; j++) {
//     answer[j] *= leftMul;
//     leftMul *= nums[j];
//   }
//   return answer;
// }

function productExceptSelf(nums: number[]): number[] {
  const answer = new Array(nums.length);

  answer[0] = 1;
  for (let i = 1; i < nums.length; i++) {
    answer[i] = answer[i - 1] * nums[i - 1];
  }

  let rightMul = 1;
  for (let i = nums.length - 1; i >= 0; i--) {
    answer[i] *= rightMul;
    rightMul *= nums[i];
  }
  return answer;
}
