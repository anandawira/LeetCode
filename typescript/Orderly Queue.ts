function orderlyQueue(s: string, k: number): string {
  if (k > 1) {
    return [...s].sort().join('');
  }

  let result = s;
  for (let i = 0; i < s.length; i++) {
    const newString = s.substr(i) + s.substr(0, i);
    if (result.localeCompare(newString) > 0) {
      result = newString;
    }
  }
  return result;
}
