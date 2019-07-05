def solve():
	f = open("in")
	fied = []
	for line in f.readlines():
		fied.append([int(x) for x in line.split()])

	ans = [[fied[0][0]]]
	for r in range(1, len(fied)):
		tmp = []
		for c in range(r + 1):
			if c == 0:
				tmp.append(ans[-1][0] + fied[r][c])
			elif c == r:
				tmp.append(ans[-1][-1] + fied[r][c])
			else:
				t = max(ans[-1][c - 1], ans[-1][c])
				tmp.append(t + fied[r][c])
		ans.append(tmp)

	return max(ans[-1])

print solve()