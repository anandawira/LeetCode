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

class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function maxProduct(root: TreeNode | null): number {
  const subtreeSums = new Set<number>();

  const dfs = (node: TreeNode | null): number => {
    if (node === null) return 0;
    const sum = node.val + dfs(node.left) + dfs(node.right);
    subtreeSums.add(sum);
    return sum;
  };

  const totalSum = dfs(root);
  let result = 0;

  subtreeSums.forEach((val) => {
    result = Math.max(result, val * (totalSum - val));
  });

  return result%(10**9+7);
}
