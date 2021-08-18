function numDecodings(s: string): number {
  const decodingDp = (data: string, k: number, memo: number[]): number => {
    if (memo[k] !== undefined) return memo[k];

    if (k === 0) return 1;

    const s = data.length - k;

    if (data[s] === '0') return 0;

    let result = decodingDp(data, k - 1, memo);

    if (k >= 2 && parseInt(data.substr(s, 2)) <= 26) {
      result += decodingDp(data, k - 2, memo);
    }

    memo[k] = result;

    return result;
  };

  const memo: number[] = new Array(s.length + 1);
  return decodingDp(s, s.length, memo);
}
