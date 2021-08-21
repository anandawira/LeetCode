/**
 Do not return anything, modify board in-place instead.
 */
function solveSudoku(board: string[][]): void {
  const isValid = (board: string[][], row: number, col: number, c: string) => {
    const blockRow = Math.floor(row / 3) * 3;
    const blockCol = Math.floor(col / 3) * 3;
    for (let i = 0; i < 9; i++) {
      if (board[row][i] === c || board[i][col] === c) return false;
      const curRow = blockRow + Math.floor(i / 3);
      const curCol = blockCol + Math.floor(i % 3);
      if (board[curRow][curCol] === c) return false;
    }
    return true;
  };

  const dfs = (board: string[][]): boolean => {
    // for every cell in the sudoku
    for (let row = 0; row < 9; row++) {
      for (let col = 0; col < 9; col++) {
        // if its empty
        if (board[row][col] !== '.') continue;
        // try every number 1-9
        for (let i = 1; i <= 9; i++) {
          const c = i.toString();
          // if that number is valid
          if (isValid(board, row, col, c)) {
            board[row][col] = c;
            // continue search for that board, ret true if solution is reached
            if (dfs(board)) return true;
          }
        }
        // solution wasnt found for any num 1-9 here, must be a dead end...
        // set the current cell back to empty
        board[row][col] = '.';
        // ret false to signal dead end
        return false;
      }
    }
    // all cells filled, must be a solution
    return true;
  };

  dfs(board);
}
