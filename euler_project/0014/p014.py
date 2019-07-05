def cal(n):
	ans = 1
	while True:
		if n == 1:
			break
		elif n % 2 == 0:
			n /= 2
		else:
			n = 3 * n + 1
		ans += 1
	return ans

assert cal(1) == 1
assert cal(13) == 10
assert cal(2) == 2

def solve(nmax):
	length = 1
	ans = 1
	for i in range(1, nmax + 1):
		tmp = cal(i)
		if length < tmp:
			ans, length = i, tmp
	return ans

print solve(1000000)