def jiecheng(n):
	ans = 1
	for i in range(1, n + 1):
		ans *= i
	return ans

def cal_comb(n, r):
	return jiecheng(n) / jiecheng(r) / jiecheng(n - r)

def solve():
	cnt = 0
	for n in range(1, 101):
		for r in range(1, n):
			tmp = cal_comb(n, r)
			if tmp > 1000000:
				cnt += 1
	return cnt

print solve()