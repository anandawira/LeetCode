/**
 * @param {string} s
 * @return {string}
 */
var reformat = function (s) {
	const alphas = [...s].filter((char) => /[a-z]/.test(char));
	const nums = [...s].filter((char) => /\d/.test(char));

	const isAlphaNow = alphas.length >= nums.length;
	let result = '';

	if (!([-1, 0, 1].includes(alphas.length-nums.length))){
		return result
	}

	i = 0;
	
	while (true) {
		const first = (isAlphaNow ? alphas : nums)[i];
		if (first == null) {
			break;
		}
		result += first;

		const second = (isAlphaNow ? nums : alphas)[i];
		if (second == null) {
			break;
		}
		result += second;

		i++;
	}

	return result;
};
reformat