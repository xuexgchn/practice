/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findErrorNums = function (nums) {
    var count = Array(100002);
    var res = [0, 0];

    for (let i = 0; i < count.length; i++) {
        count[i] = 0;
    }
    for (let i = 0; i < nums.length; i++) {
        count[nums[i]]++;
    }

    for (let i = 1; i <= nums.length; i++) {
        if (count[i] == 2) {
            res[0] = i;
        }
        if (count[i] == 0) {
            res[1] = i;
        }
    }
    return res;
};

console.log(findErrorNums([1, 2, 2, 4]));