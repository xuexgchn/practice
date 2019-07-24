/**
 * @param {number[][]} points
 * @return {number}
 */
// p = (a+b+c)/2
// s = sqrt(p(p-a)(p-b)(p-c))
var largestTriangleArea = function (points) {
    let length = points.length;
    let res = 0.0;

    for (let i = 0; i < length - 2; i++) {
        for (let j = i + 1; j < length - 1; j++) {
            for (let k = j + 1; k < length; k++) {
                let a = getLength(points[i], points[j]);
                let b = getLength(points[j], points[k]);
                let c = getLength(points[i], points[k]);

                let p = (a + b + c) / 2.0;
                let s = p * (p - a) * (p - b) * (p - c);
                if (s < 0) {
                    continue
                }
                res = Math.max(res, Math.sqrt(s));
            }
        }
    }

    return res
};

var getLength = function (pointX, pointY) {
    let tmp = (pointX[0] - pointY[0]) * (pointX[0] - pointY[0]) + (pointX[1] - pointY[1]) * (pointX[1] - pointY[1]);
    return Math.sqrt(tmp)
}

console.log(largestTriangleArea([[0, 0], [0, 1], [1, 0], [0, 2], [2, 0]]));
console.log(largestTriangleArea([[-35, 19], [40, 19], [27, -20], [35, -3], [44, 20], [22, -21], [35, 33], [-19, 42], [11, 47], [11, 37]]));