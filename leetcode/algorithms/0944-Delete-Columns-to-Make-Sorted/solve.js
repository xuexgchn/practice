/**
 * @param {string[]} A
 * @return {number}
 */
var minDeletionSize = function (A) {
    let ans = 0;

    for (let i = 0; i < A[0].length; i++) {
        for (let j = 0; j < A.length - 1; j++) {
            if (A[j].charAt(i) > A[j + 1].charAt(i)) {
                ans++;
                break;
            }
        }
    }

    return ans;
};

console.log(minDeletionSize(["cba", "daf", "ghi"]));