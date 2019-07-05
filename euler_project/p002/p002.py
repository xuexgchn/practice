def fibs(n):
	res = [0, 1]
	ans = 0
	while res[-1] <= n:
		if res[-1] % 2 == 0:
			ans += res[-1]
		res.append(res[-2] + res[-1])
	return ans

assert fibs(3) == 2
assert fibs(7) == 2
print fibs(4000000)
