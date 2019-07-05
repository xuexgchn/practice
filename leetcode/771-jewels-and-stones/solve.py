# https://leetcode-cn.com/problems/jewels-and-stones/

class Solution:
    def numJewelsInStones(self, J: str, S: str) -> int:
        sum = 0
        for s in J:
            sum += S.count(s)
        return sum