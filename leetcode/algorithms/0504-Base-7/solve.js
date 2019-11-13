/**
 * @param {number} num
 * @return {string}
 */
var convertToBase7 = function(num) {
    if (num === 0) return "0";
    
    let ans = "";
    let base = 7;
    let isPositive = true;

    if (num < 0) {
        num = -num;
        isPositive = false;
    }

    while (num > 0) {
        ans = num % base + ans;
        num = parseInt(num / 7)
    }

    return isPositive ? ans : "-" + ans;    
};

console.log(convertToBase7(100));
console.log(convertToBase7(-7));