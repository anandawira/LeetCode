/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function (n) {
	const half = Math.floor(n / 2);
	const positiveNumbers = Array.from(Array(half).keys(), x=>x+1);
	const negativeNumbers = positiveNumbers.map((x)=>x*-1)

	const result = [...positiveNumbers, ...negativeNumbers];
	if (n % 2 === 1) {
		result.push(0);
	}
	return result;
};