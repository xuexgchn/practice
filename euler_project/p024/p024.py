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
	numbers = list(range(10))
	res = perms(numbers)
	ans = ""
	for x in res[1000000 - 1]:
		ans += str(x)
	return ans

print solve()
