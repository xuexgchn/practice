def solve():
	numbers = [x * (x + 1) / 2 for x in range(1, 1000)]
	f = open("in", "r")
	cnt = 0

	for s in f.read().split(','):
		tmp = sum([ord(x) - ord('A') + 1 for x in s.strip('"')])
		if tmp in numbers:
			cnt += 1
	return cnt

print solve()


