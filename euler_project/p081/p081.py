mat = []
f = open('p081_matrix.txt', 'r')
for l in f.readlines():
	l = l.rstrip()
	digits = l.split(',')
	numbers = []
	for d in digits:
		numbers.append(int(d))
	mat.append(numbers)


initial_value = 0
initial_length = 80
ans = [[initial_value] * initial_length] * initial_length

for i in range(initial_length):
	for j in range(initial_length):
		if i == 0 and j == 0:
			ans[i][j] = mat[i][j] 
		elif i == 0 and j > 0: 
			ans[i][j] = ans[i][j - 1] + mat[i][j]
		elif j == 0 and i > 0:
			ans[i][j] = ans[i - 1][j] + mat[i][j]
		else:
			ans[i][j] = min(ans[i - 1][j], ans[i][j - 1]) + mat[i][j]

print ans[initial_length - 1][initial_length - 1]

			 
			


