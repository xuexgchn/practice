
# >>> int("03")
# 3
def turn(s):
	if s[0] == '0':
		return int(s[1])
	else:
		return int(s)

# f = open("in.txt", "r")
# grid = []
# for line in f.readlines():
# 	digits = line.rstrip("\n").split(' ')
# 	grid.append([int(x) for x in digits])
# print grid[0][2]

def solve():
	f = open("in", "r")
	s = f.read()
	ls = s.split()
	fied = []
	
	for i in range(20):
		tmp = []
		for j in range(20):
			# tmp.append(turn(ls[i * 20 + j]))
			tmp.append(int(ls[i * 20 + j]))
		fied.append(tmp)

	ans = 0

	for i in range(17):
		for j in range(20):
			tmp = fied[i][j] * fied[i + 1][j] * fied[i + 2][j] * fied[i + 3][j]
			if tmp > ans:	ans = tmp

	for i in range(20):
		for j in range(17):
			tmp = fied[i][j] * fied[i][j + 1] * fied[i][j + 2] * fied[i][j + 3]
			if tmp > ans:	ans = tmp

	for i in range(17):
		for j in range(17):
			tmp = fied[i][j] * fied[i + 1][j + 1] * fied[i + 2][j + 2] * fied[i + 3][j + 3]
			if tmp > ans:	ans = tmp

	for i in range(17):
		for j in range(3, 20):
			tmp = fied[i][j] * fied[i + 1][j - 1] * fied[i + 2][j - 2] * fied[i + 3][j - 3]
			if tmp > ans:	ans = tmp

	return ans
	
print solve()
