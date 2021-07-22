/**
 * @param {number} n
 * @return {number}
 */

var fib = function (n) {
	const phi = (1 + Math.sqrt(5)) / 2;
	return Math.round(Math.pow(phi, n) / Math.sqrt(5));
};
