# >>> 9 ** 5
# 59049
# >>> 9 ** 5 * 6
# 354294
# >>> 9 ** 5 * 7
# 413343
# >>> 

def solve():
	res = []
	for x in range(2, 9 ** 5 * 6):
		tmp = [int(i) for i in str(x)]
		ans = [t ** 5 for t in tmp]
		if x == sum(ans):
			res.append(x)

	return sum(res)

print solve()
