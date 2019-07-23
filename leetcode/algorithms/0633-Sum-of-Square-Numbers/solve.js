var judgeSquareSum = function(c) {
    if (c < 0) {
        return false
    } else if (c < 3) {
        return true
    } else {
        let l = Math.ceil(Math.sqrt(c));
        for (let i = 0; i < l; i++) {
            let j = Math.ceil(Math.sqrt(c - i * i));
            if(i <= j && i * i + j * j == c) {
                return true
            }
        }
        return false
    }
};

console.log(judgeSquareSum(5));
console.log(judgeSquareSum(3));
console.log(judgeSquareSum(8));

