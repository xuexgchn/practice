def solve():
	f = open("in", "r")
	sum = 0
	for i in range(100):
		line = f.readline()
		sum += int(line)
	return str(sum)[:10]

print solve()