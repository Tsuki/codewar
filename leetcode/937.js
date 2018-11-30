/**
 * @param {string[]} logs
 * @return {string[]}
 */
const reorderLogFiles = function (logs) {
    let llog = [];
    let dlog = [];
    for (const log of logs) {
        if (log.charCodeAt(log.indexOf(" ") + 1) >= 97) {
            llog.push(log)
        } else {
            dlog.push(log)
        }
    }
    llog.sort((a, b) => {
        a = a.slice(a.indexOf(" "), a.length);
        b = b.slice(b.indexOf(" "), b.length);
        return a.localeCompare(b)
    });
    return [...llog, ...dlog]
};

const a = ["a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"];
const b = ["g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"];
const assert = require('assert');
describe('Array', function () {
    describe('numSubarraysWithSum', function () {
        it('1', function () {
            assert.deepEqual(reorderLogFiles(a), b);
        });
    });
});