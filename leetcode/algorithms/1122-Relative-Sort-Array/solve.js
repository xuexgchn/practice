/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @return {number[]}
 */
var relativeSortArray = function (arr1, arr2) {
    var count = Array(1005);
    count.fill(0);

    for (let i = 0; i < arr1.length; i++) {
        count[arr1[i]]++;
    }

    let k = 0;
    for (let i = 0; i < arr2.length; i++) {
        for (let j = 0; j < count[arr2[i]]; j++) {
            arr1[k] = arr2[i];
            k++;
        }
        count[arr2[i]] = 0;
    }

    for (let i = 0; i < count.length; i++) {
        for (let j = 0; j < count[i]; j++) {
            arr1[k] = i;
            k++;
        }
    }

    return arr1;
};

console.log(relativeSortArray([2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19], [2, 1, 4, 3, 9, 6]));