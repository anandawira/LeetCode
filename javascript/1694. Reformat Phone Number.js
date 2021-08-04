/**
 * @param {string} number
 * @return {string}
 */
var reformatNumber = function (number) {
    number = number.replace(/[^0-9]/g, '');

    const blocks = [];

    while (number.length > 4) {
        blocks.push(number.slice(0, 3));
        number = number.slice(3);
    }

    if (number.length === 4) {
        blocks.push(number.slice(0, 2), number.slice(2));
    } else {
        blocks.push(number);
    }

    return blocks.join('-');
};
