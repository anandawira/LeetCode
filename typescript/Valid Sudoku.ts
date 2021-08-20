function isValidSudoku(board: string[][]): boolean {
  const rows: Set<string>[] = Array(9)
    .fill(null)
    .map(() => new Set());

  const columns = Array(9)
    .fill(null)
    .map(() => new Set());

  const subboxes = Array(9)
    .fill(null)
    .map(() => new Set());

  for (let rowIndex = 0; rowIndex < 9; rowIndex++) {
    for (let colIndex = 0; colIndex < 9; colIndex++) {
      const curVal = board[rowIndex][colIndex];
      if (curVal === '.') continue;

      const subboxesIndex =
        Math.floor(rowIndex / 3) * 3 + Math.floor(colIndex / 3);
      if (
        rows[rowIndex].has(curVal) ||
        columns[colIndex].has(curVal) ||
        subboxes[subboxesIndex].has(curVal)
      ) {
        return false;
      } else {
        rows[rowIndex].add(curVal);
        columns[colIndex].add(curVal);
        subboxes[subboxesIndex].add(curVal);
      }
    }
  }

  return true;
}
