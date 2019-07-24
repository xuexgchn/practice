/**
 * @param {character[]} letters
 * @param {character} target
 * @return {character}
 */
var nextGreatestLetter = function(letters, target) {
    let length = letters.length;
    let i = 0;
    let j = length - 1;
    let mid = 0;
    let res = 0;

    while (i <= j) {
        mid = parseInt((i + j) / 2);
        if (letters[mid] <= target) {
            i = mid + 1;
        } else {
            j = mid - 1;
        }
    }

    if (i > mid) {
        res = i;
    } else if (j < mid) {
        res = mid;
    } else {
        res = mid + 1;
    }

    return letters[res % length]
};

