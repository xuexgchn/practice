/**
 * @param {number[]} bills
 * @return {boolean}
 */
var lemonadeChange = function (bills) {
    if (bills.length < 1) {
        return false;
    }
    if (bills[0] != 5) {
        return false;
    }

    let count = [1, 0, 0];

    for (let i = 1; i < bills.length; i++) {
        switch (bills[i]) {
            case 5:
                count[0]++;
                break;
            case 10:
                count[1]++;
                count[0]--;

                if (count[0] < 0) {
                    return false;
                }
                break;
            case 20:
                count[2]++;
                if (count[1] > 0) {
                    count[1]--;
                } else {
                    count[0] -= 2;
                }

                count[0]--;

                if (count[0] < 0) {
                    return false;
                }
                break;
        }
    }

    return true
};


// console.log(lemonadeChange([5, 5, 5, 10, 20]));
// console.log(lemonadeChange([5, 5, 10]));
// console.log(lemonadeChange([10, 10]));
console.log(lemonadeChange([5, 5, 10, 10, 20]));