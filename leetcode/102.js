/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */


function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * @param {TreeNode} root
 * @return {number[][]}
 */

const levelOrder = function (root) {
    if (root === null) {
        return []
    }
    const result = [[root.val]];
    let rot = [root];
    while (true) {
        let [a, b] = getRoot(rot);
        rot = b;
        if (a.length > 0)
            result.push(a);
        if (rot.length === 0)
            break;
    }
    return result
};

function getRoot(root) {
    let a = [], b = [];
    for (const r of root) {
        if (r.left) {
            a.push(r.left.val);
            b.push(r.left)

        }
        if (r.right) {
            a.push(r.right.val);
            b.push(r.right)

        }
    }
    return [a.filter(c => c == null), b.filter(c => c == null)]
}

const assert = require('assert');
const order = new TreeNode(3);
order.left = new TreeNode(9);
order.right = new TreeNode(20);
order.right.left = new TreeNode(15);
order.right.right = new TreeNode(7);
describe('Array', function () {
    describe('numJewelsInStones', function () {
        it('1', function () {
            assert.deepEqual(levelOrder(order), [[3], [9, 20], [15, 7]]);
        });
    });
});
