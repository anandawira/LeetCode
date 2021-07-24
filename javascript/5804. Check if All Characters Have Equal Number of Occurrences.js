/**
 * @param {string} s
 * @return {boolean}
 */
var areOccurrencesEqual = function (s) {
	const charCount = new Map();

	[...s].forEach((c) => {
		charCount.set(c, (charCount.get(c) || 0) + 1);
	});

	const countArray = [...charCount.values()];

	for (count of countArray) {
		if (count !== countArray[0]) {
			return false;
		}
	}

	return true;
};

console.log(areOccurrencesEqual("abacbc"));
