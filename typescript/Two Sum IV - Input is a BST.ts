/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function findTarget(root: TreeNode | null, k: number): boolean {
  const nums = new Array<number>();

  // Convert binary tree to sorted array
  const treeToArray = (node: TreeNode | null): void => {
    if (!node) return;

    treeToArray(node.left);
    nums.push(node.val);
    treeToArray(node.right);
  };

  treeToArray(root);

  // Use Two Sum II algo
  let l = 0;
  let r = nums.length - 1;
  while (l < r) {
    const sum = nums[l] + nums[r];
    if (sum === k) return true;
    // sum < k ? l++ : r--;
    if (sum < k) {
      l++;
    } else {
      r--;
    }
  }
  return false;
}
