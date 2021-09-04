function sumOfDistancesInTree(n: number, edges: number[][]): number[] {
  const tree = Array(n)
    .fill(null)
    .map(() => new Set<number>());

  const result = Array(n).fill(0);
  const count = Array(n).fill(1);

  edges.forEach((edge) => {
    tree[edge[0]].add(edge[1]);
    tree[edge[1]].add(edge[0]);
  });

  const dfs = (root: number, pre: number) => {
    tree[root].forEach((i) => {
      if (i !== pre) {
        dfs(i, root);
        count[root] += count[i];
        result[root] += result[i] + count[i];
      }
    });
  };

  const dfs2 = (root: number, pre: number) => {
    tree[root].forEach((i) => {
      if (i !== pre) {
        result[i] = result[root] - count[i] + n - count[i];
        dfs2(i, root);
      }
    });
  };

  dfs(0, -1);
  dfs2(0, -1);
  return result;
}
