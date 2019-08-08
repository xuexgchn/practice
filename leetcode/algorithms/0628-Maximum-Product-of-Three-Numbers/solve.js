/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumProduct = function (nums) {
    nums.sort((a, b) => a - b);

    var len = nums.length;
    return Math.max(nums[len - 1] * nums[len - 2] * nums[len - 3], nums[0] * nums[1] * nums[len - 1]);
};


console.log(maximumProduct([1, 2, 3]));
console.log(maximumProduct([1, 2, 3, 4]));