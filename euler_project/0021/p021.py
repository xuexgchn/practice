def d(x):
	return sum([i for i in range(1, x) if x % i == 0])

def solve(nmax):
	nums = [d(x) for x in range(nmax + 1)]
	ans = []
	for i in range(nmax + 1):
		if i == d(nums[i]) and nums[i] != i:
			ans.append(i)
	return sum(ans)

print solve(10000)