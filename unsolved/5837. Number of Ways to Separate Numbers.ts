function numberOfCombinations(num: string): number {
  const dp = (
    data: string,
    k: number,
    lastNum: number,
    lastDigits: number,
  ): number => {
    if (k === 0) return 1;

    const s = data.length - k;

    if (data[s] === '0') return 0;

    let result = 0;

    for (let digits = lastDigits; digits <= k; digits++) {
      const num = parseInt(data.substr(s, digits));
      if (num >= lastNum) {
        result += dp(data, k - digits, num, digits);
      }
    }
    return result;
  };
  return dp(num, num.length, 0, 1);
}
