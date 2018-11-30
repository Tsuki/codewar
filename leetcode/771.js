/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
const numJewelsInStones = function (J, S) {
    let sum = 0;
    for (const j of J) {
        for (const s of S) {
            if (j === s) {
                sum++;
            }
        }
    }
    return sum
};

const assert = require('assert');
describe('Array', function () {
    describe('numJewelsInStones', function () {
        it('1', function () {
            assert.equal(numJewelsInStones("aA", "aAAbbbb"), 3);
        });
    });
});
