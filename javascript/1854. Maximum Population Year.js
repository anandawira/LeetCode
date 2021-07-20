/**
 * @param {number[][]} logs
 * @return {number}
 */
var maximumPopulation = function (logs) {
	// const getYears = () => {
	// 	const [birth, death] = log;
	// 	let result = [...Array(death - birth).keys()].map((val) => val + birth);
	// 	return result;
	// };

	// years = [];
	// for (log of logs) {
	// 	years = [...years, ...getYears(log)]

	// }

	// return years;

	logs.sort((a, b) => a[0] - b[0]);

	const years = new Map();

	for (log of logs) {
		const [birth, death] = log;
		for (let i = 0; i < death - birth; i++) {
			const year = i + birth;
			const count = (years.get(year) | 0) + 1;
			years.set(year, count);
		}
	}

	let maxYear = 1950;
	let maxCount = 0;

	years.forEach((count, year) => {
		if (count > maxCount) {
			maxYear = year;
			maxCount = count;
		}
	});

	return maxYear;
};

console.log(
	maximumPopulation([
		[1950, 1961],
		[1960, 1971],
		[1970, 1981],
	]),
);
