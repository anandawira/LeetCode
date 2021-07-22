/**
 * @param {number} low
 * @param {number} high
 * @return {number[]}
 */
var sequentialDigits = function (low, high) {
	const lowDigit = Math.ceil(Math.log10(low));
	const highDigit = Math.ceil(Math.log10(high));

	const getFullNumber = (start, digit) => {
		if (digit === 1) {
			return start;
		}
		return getFullNumber(start, digit - 1) * 10 + (start + digit - 1);
	};

	const digits = [];

	for (let d = lowDigit; d <= highDigit; d++) {
		for (let start = 1; start <= 10 - d; start++) {
			digits.push(getFullNumber(start, d));
		}
	}

	return digits
		.filter((val) => val >= low && val <= high)
		.sort((a, b) => a > b);
};
