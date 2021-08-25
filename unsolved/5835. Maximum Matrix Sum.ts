function maxMatrixSum(matrix: number[][]): number {
  let negativeCount = 0;
  let absSum = 0;
  let absMin = Math.abs(matrix[0][0]);

  matrix.forEach((row) => {
    row.forEach((cell) => {
      if (cell < 0) negativeCount++;

      const absVal = Math.abs(cell);
      absSum += absVal;
      absMin = Math.min(absMin, absVal);
    });
  });

  if (negativeCount % 2 === 0) {
    return absSum;
  } else {
    return absSum - absMin * 2;
  }
}
