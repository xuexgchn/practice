/**
 * @param {number} n
 * @return {number}
 */
var tribonacci = function(n) {
    var res = Array(38);
    res[0] = 0;
    res[1] = 1;
    res[2] = 1;

    for (let i = 3; i <= n; i++) {
        res[i] = res[i - 1] + res[i - 2] + res[i - 3];
    }

    return res[n];
};

console.log(tribonacci(2));
console.log(tribonacci(4));
console.log(tribonacci(25));