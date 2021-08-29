// function kthLargestNumber(nums: string[], k: number): string {
//   const topK = new Array();
//   for (let i = 0; i < k; i++) {
//     topK.push(parseInt(nums[i]));
//   }

//   let min = Math.min(...topK);
//   for (let i = k; i < nums.length; i++) {
//     const num = parseInt(nums[i]);
//     if (num > min) {
//     }
//   }
// }

function kthLargestNumber(nums: string[], k: number): string {
  nums.sort((a, b) => {
    if (b.length !== a.length) return b.length - a.length;
    for (let i = 0; i < a.length; i++) {
      if (b[i] !== a[i]) return parseInt(b[i]) - parseInt(a[i]);
    }
    return 0;
  });

  return nums[k - 1];
}
