/**
 * @param {string} S
 * @return {string}
 */
var removeDuplicates = function(S) {
    let stack = [];
    
    for (let i = 0; i < S.length; i++) {
        let tmp = stack.pop();
        if (S[i] == tmp) {
            continue;
        } else {
            stack.push(tmp);
            stack.push(S[i]);
        }
    }

    return stack.join("");
};

console.log(removeDuplicates("abbaca"));