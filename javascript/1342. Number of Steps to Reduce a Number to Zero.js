/**
 * @param {number} num
 * @return {number}
 */
var numberOfSteps = function (num) {
	let count = 0;

	while (num !== 0) {
		count++;
		if (num % 2 === 0) {
			num >>= 1;
		} else {
			num--;
		}
	}

	return count;
};

numberOfSteps(14);
