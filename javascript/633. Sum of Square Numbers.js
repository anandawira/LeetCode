/**
 * @param {number} c
 * @return {boolean}
 */
var judgeSquareSum = function (c) {
  if (Number.isInteger(Math.sqrt(c))) return true;

  for (let a = 0; a * a <= c; a++) {
    if (Number.isInteger(Math.sqrt(c - a * a))) return true;
  }
  return false;
};
