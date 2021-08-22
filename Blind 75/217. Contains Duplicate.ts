// Convert to set
function containsDuplicate(nums: number[]): boolean {
  return nums.length > new Set([...nums]).size;
}

// Using set has
// function containsDuplicate(nums: number[]): boolean {
//   const prevs = new Set();
//   for (const n of nums) {
//     if (prevs.has(n)) return true;
//     prevs.add(n);
//   }
//   return false;
// }
