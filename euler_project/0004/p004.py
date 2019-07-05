# def is_palindrome(n):
#        str(n) == str(n)[::-1]

def solve():	
	ans = 0;
	for x in range(100, 1000):
		for y in range(100, 1000):
			s = str(x * y);
			ls = list(s)
			ls.reverse()
			if (list(s) == ls and x * y > ans):
				ans = x * y
	return ans

print solve()
