function numberOfUniqueGoodSubsequences(binary: string): number {
  const subs = new Set();
  const dp = (binary: string, subseq: string, index: number) => {
    const cur = subseq + binary[index];
    subs.add(cur);

    if (index !== binary.length - 1) {
      dp(binary, '', index + 1);

      if (cur !== '0') {
        dp(binary, cur, index + 1);
      }
    }
  };

  dp(binary, '', 0);

  return subs.size;
}
