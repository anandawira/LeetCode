function maxArea(height: number[]): number {
  let l = 0;
  let r = height.length - 1;
  let res = 0;
  while (l < r) {
    const minHeight = Math.min(height[l], height[r]);
    const area = minHeight * (r - l);

    res = Math.max(res, area);

    if (height[l] < height[r]) {
      l++;
    } else {
      r--;
    }
  }

  return res;
}
