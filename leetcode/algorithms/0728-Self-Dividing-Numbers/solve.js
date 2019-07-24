var selfDividingNumbers = function(left, right) {
    let res = [];

    for (let i = left; i <= right; i++) {
        
        let tmp = i;
        let tmp2 = 1;
        while (tmp != 0) {
            tmp2 = tmp % 10;
            if (tmp2 == 0 || i % tmp2 != 0) {
                break;
            } else {
                tmp = parseInt(tmp / 10);
            }
        }
        if (tmp == 0) {
            res.push(i);
        }
    }

    return res
};

console.log(selfDividingNumbers(1, 22));