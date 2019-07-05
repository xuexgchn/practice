# https://leetcode-cn.com/problems/sort-characters-by-frequency/

import queue 
class Solution:
    def frequencySort(self, s: str) -> str:
        dic = {}
        for i in s:
            if i not in dic:
                dic[i] = 1
            else:
                dic[i] = dic[i] + 1

        que = queue.PriorityQueue()
        for k in dic:
            que.put([dic[k], k])
        res = ""
        while que.qsize() > 0:
            s = que.get()
            res += s[1] * s[0]
        return res[::-1]

assert Solution().frequencySort("tree") == "eert" or Solution().frequencySort("tree") == "eetr"
assert Solution().frequencySort("cccaaa") == "cccaaa"
assert Solution().frequencySort("Aabb") == "bbAa" or Solution().frequencySort("Aabb") == "bbaA"