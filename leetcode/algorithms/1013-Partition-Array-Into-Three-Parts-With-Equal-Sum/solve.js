/**
 * @param {number[]} A
 * @return {boolean}
 */
var canThreePartsEqualSum = function (A) {
    var sum = 0;

    for (let i = 0; i < A.length; i++) {
        sum += A[i];
    }
    if (sum % 3 != 0) {
        return false
    }

    var targetA = sum / 3;
    var targetB = sum / 3;

    for (let i = 0; i < A.length; i++) {
        if (targetA != 0) {
            targetA -= A[i];
        } else if (targetB != 0) {
            targetB -= A[i];
        }
        if (targetB === 0 && targetA === 0) {
            return true
        }
    }

    return false
};

// console.log(canThreePartsEqualSum([0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1]));
// console.log(canThreePartsEqualSum([0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1]));
// console.log(canThreePartsEqualSum([3, 3, 6, 5, -2, 2, 5, 1, -9, 4]));

console.log(canThreePartsEqualSum([18, 12, -18, 18, -19, -1, 10, 10]));