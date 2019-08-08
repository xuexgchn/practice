/**
 * @param {number} n
 * @return {number}
 */
var arrangeCoins = function(n) {
    var i = 0;
    var sum = 0;

    while (sum <= n) {
        sum += i;
        i++;
    }

    return i - 2;
};

console.log(arrangeCoins(5));
console.log(arrangeCoins(8));
console.log(arrangeCoins(10));