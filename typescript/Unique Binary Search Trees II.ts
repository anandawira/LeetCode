/**
 * Definition for a binary tree node.
class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.left = (left===undefined ? null : left)
        this.right = (right===undefined ? null : right)
    }
}
 */

function generateTrees(n: number): Array<TreeNode | null> {
  const memo = new Map();
  const dfs = (start: number, end: number): Array<TreeNode> | Array<null> => {
    if (memo.has(`${start},${end}`)) {
      return memo.get(`${start},${end}`);
    }

    if (start > end) return [null];

    const result = new Array();
    for (let i = start; i <= end; i++) {
      const leftTrees = dfs(start, i - 1);
      const rightTrees = dfs(i + 1, end);

      for (let left of leftTrees) {
        for (let right of rightTrees) {
          const tree = new TreeNode(i, left, right);
          result.push(tree);
        }
      }
    }
    memo.set(`${start},${end}`, result);
    return result;
  };

  return dfs(1, n);
}
