# https://leetcode-cn.com/problems/baseball-game/

class Solution:
    def calPoints(self, ops: list) -> int:
        ans = []
        for op in ops:
            if op == "C":
                ans.pop()
            elif op == "D":
                tmp = ans[-1]
                ans.append(tmp * 2)
            elif op == "+":
                tmp1 = ans[-1]
                tmp2 = ans[-2]
                ans.append(tmp1 + tmp2)
            else:
                ans.append(int(op))

        return sum(ans)

assert Solution().calPoints(["5","2","C","D","+"]) == 30
assert Solution().calPoints(["5","-2","4","C","D","9","+","+"]) == 27