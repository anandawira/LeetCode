function minTimeToType(word: string): number {
  let result = word.length;
  for (let i = 0; i < word.length; i++) {
    const before = i === 0 ? 97 : word.charCodeAt(i - 1);
    const current = word.charCodeAt(i);

    result += Math.min(
      Math.abs(before - current),
      Math.abs(Math.abs(before - current) - 26),
    );
  }
  return result;
}

minTimeToType('abc');
