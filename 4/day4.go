package main

import (
    "bufio"
		"fmt"
		"strings"
		"regexp"
		"strconv"
    "os"
)

type matcher func(x string) bool

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func byr (s string) bool {
	v,_ := strconv.Atoi(s)
	return 1920 <= v && v <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func iyr (s string) bool {
	v,_ := strconv.Atoi(s)
	return 2010 <= v && v <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func eyr (s string) bool {
	v,_ := strconv.Atoi(s)
	return 2020 <= v && v <= 2030
}

//hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
func hgt (s string) bool {
	var leadingInt = regexp.MustCompile(`^[-+]?\d+`)
	var hgtType = regexp.MustCompile(`cm|in`)
	v,_ := strconv.Atoi(leadingInt.FindString(s))
	switch (hgtType.FindString(s)) {
		case "cm": return 150 <= v && v <= 193
		case "in": return 59 <= v && v <= 76
	}
	return false
}

//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hcl (s string) bool {
	matched, _ := regexp.MatchString(`#[0-9a-f]{6}`, s)
	return matched
}
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func ecl (s string) bool {
	matched, _ := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, s)
	return matched
}
//pid (Passport ID) - a nine-digit number, including leading zeroes.
func pid (s string) bool {
	matched, _ := regexp.MatchString(`^\d{9}$`, s)
	return matched && len(s) == 9
}
//cid (Country ID) - ignored, missing or not.

func mkMap() map[string]matcher {
	m := make(map[string]matcher)
	m["byr"] = byr
	m["iyr"] = iyr
	m["eyr"] = eyr
	m["hgt"] = hgt
	m["hcl"] = hcl
	m["ecl"] = ecl
	m["pid"] = pid
	return m
} 
 
func main() {
	sum1 := 0
	fldi := mkMap()
	scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		if (len(l) == 0) {
			// reset
			// fmt.Println(fldi)
			fldi = mkMap()
		} else {
			for _,d := range l {
				p := strings.Split(d, ":")
				k := p[0]
				f,exists := fldi[k]
				if (exists && f(p[1])) {
					delete(fldi, k);
				}
			}
			
			if (len(fldi) == 0) {
				sum1 ++
				fldi = mkMap()
			}
		}
  }

    if scanner.Err() != nil {
      // handle error.
		}
		fmt.Printf("Sum1: %d\n", sum1)
  }