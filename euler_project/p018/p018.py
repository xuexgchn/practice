def solve():
	f = open("in", "r")
	grid = []
	for line in f.readlines():
		digits = line.rstrip("\n").split(' ')
		grid.append([int(x) for x in digits])

	length = len(grid)
	ans = [grid[0]]
	for r in range(1, length):
		tmp = []
		for c in range(r + 1):
			if c == 0:
				tmp.append(ans[r - 1][c] + grid[r][c])
			elif c == r:
				tmp.append(ans[r - 1][c - 1] + grid[r][c])
			else:
				tmp.append(max(ans[r - 1][c - 1], ans[r-1][c]) + grid[r][c])
		ans.append(tmp)

	return max(ans[length - 1])
print solve()