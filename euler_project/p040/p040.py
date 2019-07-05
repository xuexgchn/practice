def solve():
	ls = [x for x in range(1000000)]
	ans = ""
	res = 1
	for x in ls:
		ans += str(x)
	ls = [10 ** x for x in range(7)]
	for x in ls:
		res *= int(ans[x])

	return res

print solve()