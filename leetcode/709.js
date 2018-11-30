/**
 * @param {string} str
 * @return {string}
 */
const toLowerCase = function (str) {
    let result = "";
    for (let i = 0; i < str.length; i++) {
        const c = str.charCodeAt(i);
        if (c >= 65 && c <= 90) {
            result += String.fromCharCode(c + 32);
        } else {
            result += str[i];
        }
    }
    return result
};

const assert = require('assert');
describe('Array', function () {
    describe('numSubarraysWithSum', function () {
        it('1', function () {
            assert.equal(toLowerCase("Hello"), "hello");
        });
    });
});