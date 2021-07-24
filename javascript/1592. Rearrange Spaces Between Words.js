/**
 * @param {string} text
 * @return {string}
 */
var reorderSpaces = function (text) {
	const words = text.match(/[\S]+/g);
	const spaceCount = text.match(/ /g) ? text.match(/ /g).length : 0;
	const numSpaces = Math.floor(spaceCount / (words.length - 1));
	const spaces = ' '.repeat(isFinite(numSpaces) ? numSpaces : 0);
	const lastSpaces = ' '.repeat(
		isFinite(numSpaces) ? spaceCount % (words.length - 1) : spaceCount,
	);

	let result = '';
	words.forEach((val, index) => {
		result += val;
		if (index !== words.length - 1) {
			result += spaces;
		} else {
			result += lastSpaces;
		}
	});
	return result;
};

reorderSpaces('  hello');
