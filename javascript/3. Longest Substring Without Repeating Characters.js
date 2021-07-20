/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
	max = 0;
	substring = [];
	[...s].forEach((c) => {
		if (substring.includes(c)) {
			max = Math.max(max, substring.length);
			idx = substring.indexOf(c);

			substring.splice(0, idx + 1);
		}

		substring.push(c);
	});
	max = Math.max(max, substring.length);

	return max;
};

console.log(lengthOfLongestSubstring('dvdf'));
