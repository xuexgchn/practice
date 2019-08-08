/**
 * @param {number} m
 * @param {number} n
 * @param {number[][]} ops
 * @return {number}
 */
var maxCount = function(m, n, ops) {
    var maxM = m;
    var maxN = n;
    
    for (let i  = 0; i < ops.length; i++) {
        if (ops[i][0] < maxM) {
            maxM = ops[i][0];
        }
        if (ops[i][1] < maxN) {
            maxN = ops[i][1];
        }
    }
    return maxM * maxN;
};

console.log(maxCount(3, 3, [[2, 2], [3, 3]]));