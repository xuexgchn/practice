def factorial(n):
	ans = 1
	for i in range(1, n + 1):
		ans *= i
	return ans

def comb(m, n):
	return factorial(m) / factorial(n) / factorial(m - n)

assert comb(4, 2) == 6
assert comb(3, 2) == 3
assert comb(1, 1) == 1

print comb(40, 20)