def perms(numbers):
	if len(numbers) <= 1:
		return [numbers]
	res = []
	for i in range(len(numbers)): 
		tmp = numbers[:i] + numbers[i + 1:]
		p = perms(tmp)
		for number in p:
			res.append(numbers[i: i + 1] + number)
	return res

def solve():
	prime = [0, 2, 3, 5, 7, 11, 13, 17]

	def check(s):
		for i in range(1, 8):
			tmp = int(s[i:i + 3])
			if tmp % prime[i]:
				return False
		return True

	perm = perms([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
	sum = 0
	for x in perm:
		if x[0] == 0:
			continue
		else:
			tmp = ""
			for i in x:
				tmp += str(i)
			if check(tmp):
				sum += int(tmp)

	return sum

print solve()
	

