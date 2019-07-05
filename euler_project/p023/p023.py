def abundant_numbers():
	ans = []
	for x in range(1, 28123 + 1):
		if x < sum([i for i in range(1, x) if x % i == 0]):
			ans.append(x)
	return ans

def solve():
	ls = abundant_numbers()
	tmp = []
	length = len(ls)
	for i in range(0, length):
		for j in range(i, length):
			if ls[i] + ls[j] <= 28123:
				tmp.append(ls[i] + ls[j])
	tmp = set(tmp)
	numbers = [x for x in range(1, 28123 + 1)]
	return sum(numbers) - sum(tmp)

print solve()