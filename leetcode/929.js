/**
 * @param {string[]} emails
 * @return {number}
 */
const numUniqueEmails = function (emails) {
    const result = [];
    for (let email of emails) {
        result.push(mashMail(email));
    }
    return result.filter((item, pos) => result.indexOf(item) === pos).length
};

/**
 *
 * @param {string} email
 * @returns {string}
 */
function mashMail(email) {
    let [name, domain] = email.split("@");
    name = name.split("+");
    name = name[0].split(".");
    return `${name.join("")}@${domain}`
}

const assert = require('assert');
const order = ["test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"];

describe('Array', function () {
    describe('numJewelsInStones', function () {
        it('1', function () {
            assert.equal(numUniqueEmails(order), 2);
        });
    });
});