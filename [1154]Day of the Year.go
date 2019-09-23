//Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year. 
//
// 
// Example 1: 
//
// 
//Input: date = "2019-01-09"
//Output: 9
//Explanation: Given date is the 9th day of the year in 2019.
// 
//
// Example 2: 
//
// 
//Input: date = "2019-02-10"
//Output: 41
// 
//
// Example 3: 
//
// 
//Input: date = "2003-03-01"
//Output: 60
// 
//
// Example 4: 
//
// 
//Input: date = "2004-03-01"
//Output: 61
// 
//
// 
// Constraints: 
//
// 
// date.length == 10 
// date[4] == date[7] == '-', and all other date[i]'s are digits 
// date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019. 
//

func dayOfYear(date string) int {
	dm := map[int]int{
		1: 31,
		2: 28,
		3: 31,
		4: 30,
		5: 31,
		6: 30,
		7: 31,
		8: 31,
		9: 30,
		10: 31,
		11: 30,
		12: 31,
	}

	res := strings.Split(date, "-")
	year, _ := strconv.Atoi(res[0])
	month, _ := strconv.Atoi(res[1])
	day, _ := strconv.Atoi(res[2])

	if month == 1 {
		return day
	} else if month == 2 {
		return 31 + day
	} else {
		ress := 0
		for i := 0; i < month; i++ {
			ress += dm[i]
		}

		if year % 4 == 0 && year % 100 != 0 {
			ress++
		}
		return ress + day
	}

}