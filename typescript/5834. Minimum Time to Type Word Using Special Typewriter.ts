function minTimeToType(word: string): number {
  let result = word.length;

  let before = 97;
  for (let i = 0; i < word.length; i++) {
    const current = word.charCodeAt(i);

    result += Math.min(
      Math.abs(before - current),
      Math.abs(Math.abs(before - current) - 26),
    );

    before = current;
  }
  return result;
}

minTimeToType('abc');
