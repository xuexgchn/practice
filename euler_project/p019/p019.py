def is_leap(year):
	if ((year % 4 == 0 and year % 100 != 0) or year % 400 == 0):
		return [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
	else:
		return [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

def solve():
	cnt = 0
	rest = 0
	for year in range(1900, 2001):
		months = is_leap(year)
		for month in months:
			rest += month % 7
			rest = rest % 7 
			if year != 1900 and rest == 6:
				cnt += 1
	return cnt

print solve()
