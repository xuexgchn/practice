# https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/

class Solution:
    def minAddToMakeValid(self, S: str) -> int:
        ans = []
        for s in S:
            if s == "(":
                ans.append(s)
            elif s == ")":
                if ans == []:
                    ans.append(s)
                elif ans[-1] == "(":
                    ans.pop()
                else:
                    ans.append(s)
        return len(ans)

assert Solution().minAddToMakeValid("())") == 1
assert Solution().minAddToMakeValid("(((") == 3
assert Solution().minAddToMakeValid("()") == 0
assert Solution().minAddToMakeValid("()))((") == 4
