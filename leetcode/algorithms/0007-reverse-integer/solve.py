class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        n = x if x > 0 else -x
        res = 0
        while n:
            res = res * 10 + n % 10
            n = n // 10
        if res > 0x7fffffff:
            return 0
        return res if x > 0 else -res

assert Solution().reverse(123) == 321
assert Solution().reverse(-123) == -321
assert Solution().reverse(120) == 21
