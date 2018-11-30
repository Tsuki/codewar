function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/*
 * @param {TreeNode} root
 * @param {number} L
 * @param {number} R
 * @return {number}
 */
const rangeSumBST = function (root, L, R) {
    if (!root.val) {
        return 0
    }

    function search(root) {
        if (!root) {
            return 0
        }
        const a = search(root.left);
        const b = search(root.right);
        if (root.val >= L && root.val <= R) {
            return root.val + a + b
        } else {
            return a + b
        }
    }

    return search(root)
};

const root = new TreeNode(10);
root.left = new TreeNode(5);
root.left.left = new TreeNode(3);
root.left.right = new TreeNode(7);
root.right = new TreeNode(15);
root.right.right = new TreeNode(18);
const assert = require('assert');
describe('Array', function () {
    describe('numSubarraysWithSum', function () {
        it('1', function () {
            assert.equal(rangeSumBST(root, 7, 15), 32);
        });
    });
});