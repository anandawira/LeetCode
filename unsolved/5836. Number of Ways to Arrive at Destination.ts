function countPaths(n: number, roads: number[][]): number {
  const target = n - 1;
  let minTime = 10 ** 9;
  let waysCount = 0;

  const dp = (prevs: Set<number>, cur: number, totalTime: number) => {
    if (cur === target && totalTime <= minTime) {
      if (totalTime < minTime) {
        waysCount = 0;
        minTime = totalTime;
      }
      console.log(prevs);
      waysCount++;
    }

    roads.forEach((val) => {
      if (val[0] === cur) {
        if (!prevs.has(val[1])) {
          const newPrevs = new Set(prevs);
          newPrevs.add(val[1]);
          dp(newPrevs, val[1], totalTime + val[2]);
        }
      }
    });
  };

  dp(new Set([0]), 0, 0);
  return waysCount;
}
