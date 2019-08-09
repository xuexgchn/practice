/**
 * @param {string} S
 * @return {number[]}
 */
var diStringMatch = function(S) {
    let ans = Array(S.length + 1);

    let min = 0;
    let max = S.length;

    for (let i = 0; i < S.length; i++) {
        if (S.charAt(i) == "I") {
            ans[i] = min;
            min++;
        } else {
            ans[i] = max;
            max--;
        }
    }

    ans[S.length] = min;

    return ans;
};

console.log(diStringMatch("IDID"));