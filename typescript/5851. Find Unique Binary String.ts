function findDifferentBinaryString(nums: string[]): string {
  const combinations = new Array<string>();

  const recursive = (num: string) => {
    if (num.length === nums.length) {
      combinations.push(num);
    } else {
      recursive(num + 0);
      recursive(num + 1);
    }
  };

  recursive('');

  for (let com of combinations) {
    if (nums.indexOf(com) === -1) {
      return com;
    }
  }

  // console.log(combinations);
  return 'not found';
}
