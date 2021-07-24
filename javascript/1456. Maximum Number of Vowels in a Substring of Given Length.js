/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
 var maxVowels = function (s, k) {
	const vowels = new Set(['a', 'i', 'u', 'e', 'o']);

	const isVowel = [...s].map((x) => vowels.has(x));

	let maxCount = 0;
	let vowelCount = 0;

	isVowel.forEach((val, idx) => {
		if (val) {
			vowelCount++;
		}

		if (idx >= k) {
			if (isVowel[idx - k]) {
				vowelCount--;
			}
		}

		if (vowelCount > maxCount) {
			maxCount = vowelCount;
		}
	});

	return maxCount;
};
