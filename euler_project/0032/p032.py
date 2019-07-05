import math

def cal(n):
	factors = []
	for i in range(2, int(math.ceil(math.sqrt(n)))):
		if n % i == 0:
			factors.append([i, n / i])
	return factors

def solve():
	ans = []
	ls = set([str(i) for i in range(1, 10)])

	for n in range(1000, 10000):
		factors = cal(n)
		for factor in factors:
			tmp = ''.join(str(n))
			tmp = tmp + ''.join(str(factor[0]))
			tmp = tmp + ''.join(str(factor[1]))
			if set(tmp) == set(ls):
				ans.append(n)
	return sum(set(ans))

print solve()

