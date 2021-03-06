package main

import (
	"fmt"
	"testing"
)

func TestDay4(t *testing.T) {
	tests := []struct {
		input     []string
		expected1 int
		expected2 int
	}{
		{
			input:     []string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd", "byr:1937 iyr:2017 cid:147 hgt:183cm", "", "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884", "hcl:#cfa07d byr:1929", "", "hcl:#ae17e1 iyr:2013", "eyr:2024", "ecl:brn pid:760753108 byr:1931", "hgt:179cm", "", "hcl:#cfa07d eyr:2025 pid:166559648", "iyr:2011 ecl:brn hgt:59in"},
			expected1: 2,
			expected2: 2,
		},
		{
			input:     []string{"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980", "hcl:#623a2f", "", "eyr:2029 ecl:blu cid:129 byr:1989", "iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", "", "hcl:#888785", "hgt:164cm byr:2001 iyr:2015 cid:88", "pid:545766238 ecl:hzl", "eyr:2022", "", "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"},
			expected1: 4,
			expected2: 4,
		},
		{
			input:     []string{"eyr:1972 cid:100", "hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926", "", "iyr:2019", "hcl:#602927 eyr:1967 hgt:170cm", "ecl:grn pid:012533040 byr:1946", "", "hcl:dab227 iyr:2012", "ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277", "", "hgt:59cm ecl:zzz", "eyr:2038 hcl:74454a iyr:2023", "pid:3556412378 byr:2007"},
			expected1: 4,
			expected2: 0,
		},
		{
			input:     []string{"eyr:1972 cid:100", "hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926", "", "iyr:2019", "hcl:#602927 eyr:1967 hgt:170cm", "ecl:grn pid:012533040 byr:1946", "", "hcl:dab227 iyr:2012", "ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277", "", "hgt:59cm ecl:zzz", "eyr:2038 hcl:74454a iyr:2023", "pid:3556412378 byr:2007", "", "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980", "hcl:#623a2f", "", "eyr:2029 ecl:blu cid:129 byr:1989", "iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", "", "hcl:#888785", "hgt:164cm byr:2001 iyr:2015 cid:88", "pid:545766238 ecl:hzl", "eyr:2022", "", "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"},
			expected1: 8,
			expected2: 4,
		},
		{
			input: []string{
				"hgt:161cm eyr:2027 ecl:grn iyr:2011 hcl:#a97842 byr:1977 pid:910468396", "",
				"cid:257 ecl:gry hgt:186cm iyr:2012 byr:1941 eyr:2029 pid:108935675 hcl:#cfa07d", "",
				"eyr:2020 cid:105 iyr:2012 pid:947726115 hcl:#ceb3a1 ecl:grn byr:1966 hgt:151cm", "",
				"hcl:#888785 eyr:2027 ecl:hzl byr:1966 pid:853607760 iyr:2012 hgt:155cm", "",
				"hcl:#a97842 hgt:191cm eyr:2025 ecl:gry byr:1923 pid:574171850 iyr:2019", "",
				"byr:1955 cid:309 hcl:#a97842 pid:740105085 iyr:2020 hgt:188cm ecl:oth eyr:2029", "",
				"iyr:2016 hcl:#cfa07d eyr:2026 hgt:151cm pid:394185014 ecl:grn byr:1974", "",
				"pid:226566060 ecl:blu cid:272 hgt:188cm hcl:#efcc98 eyr:2029 iyr:2014 byr:1956", "",
				"hgt:193cm eyr:2029 pid:141707808 byr:1997 cid:83 iyr:2019 ecl:hzl hcl:#cfa07d", "",
				"iyr:2019 pid:681586971 hcl:#6b5442 hgt:165cm eyr:2022 ecl:brn byr:1985", "",
				"iyr:2013 ecl:grn pid:341584587 eyr:2027 hgt:185cm hcl:#18171d byr:1935 cid:113", "",
				"pid:175337478 ecl:oth hgt:173cm hcl:#733820 eyr:2025 byr:1960 cid:283 iyr:2018", "",
				"byr:1959 hcl:#341e13 eyr:2023 pid:566612260 hgt:176cm iyr:2017 ecl:grn", "",
				"cid:321 pid:355095309 byr:1945 hgt:161cm eyr:2029 iyr:2017 ecl:brn hcl:#733820", "",
				"eyr:2021 pid:893047101 iyr:2016 ecl:hzl hcl:#866857 byr:1988 hgt:166cm", "",
				"hcl:#7d3b0c ecl:blu pid:085336099 eyr:2024 iyr:2019 hgt:178cm byr:1999", "",
				"pid:677187953 eyr:2025 iyr:2020 hgt:163cm byr:1957 ecl:grn hcl:#cfa07d", "",
				"cid:213 byr:1987 pid:113078018 ecl:blu iyr:2013 eyr:2022 hcl:#7d3b0c hgt:157cm", "",
				"hcl:#602927 ecl:blu hgt:173cm byr:1974 pid:116377061 cid:294 eyr:2030 iyr:2010", "",
				"hgt:178cm eyr:2027 hcl:#733820 ecl:grn iyr:2014 pid:575371227 byr:1965", "",
				"hcl:#8e1608 pid:554618249 iyr:2010 hgt:176cm cid:220 ecl:brn byr:1928 eyr:2029", "",
				"eyr:2030 ecl:oth cid:177 hcl:#602927 iyr:2010 hgt:66in pid:915661465 byr:1992", "",
				"ecl:brn pid:558826437 hgt:151cm byr:1936 hcl:#fffffd eyr:2021 iyr:2012", "",
				"ecl:hzl byr:1943 iyr:2020 hgt:175cm pid:830628564 hcl:#7d3b0c eyr:2021", "",
				"hgt:182cm byr:1951 cid:175 eyr:2021 pid:635966127 ecl:blu iyr:2014 hcl:#18171d", "",
				"hcl:#733820 iyr:2011 pid:581100835 eyr:2022 ecl:grn byr:1985 hgt:192cm", "",
				"iyr:2013 ecl:grn hgt:185cm hcl:#a97842 byr:1981 eyr:2029 pid:711625030", "",
				"hgt:159cm iyr:2018 pid:610521467 eyr:2028 ecl:amb byr:1934 hcl:#602927", "",
				"ecl:blu hcl:#09d9a5 hgt:162cm iyr:2020 eyr:2025 byr:1971 pid:406714780", "",
				"hgt:179cm eyr:2022 hcl:#18171d ecl:blu pid:314891131 iyr:2015 byr:2002", "",
				"hcl:#623a2f hgt:181cm pid:442693333 byr:1990 ecl:grn eyr:2027 iyr:2011", "",
				"ecl:grn byr:1948 hgt:174cm pid:438876745 cid:321 iyr:2018 hcl:#866857 eyr:2023", "",
				"ecl:amb cid:235 byr:1938 pid:315825546 hcl:#ceb3a1 eyr:2029 iyr:2013 hgt:171cm", "",
				"hcl:#733820 byr:1988 pid:558453117 hgt:66in ecl:oth iyr:2010 eyr:2021", "",
				"byr:1926 pid:796557821 cid:155 hcl:#efcc98 hgt:159cm eyr:2023 ecl:oth iyr:2016", "",
				"hgt:155cm byr:1941 eyr:2024 cid:164 hcl:#602927 iyr:2013 pid:531781358 ecl:amb", "",
				"byr:1991 ecl:grn iyr:2020 pid:9921729009 eyr:2029 hcl:#623a2f hgt:62in", "",
				"iyr:2017 ecl:hzl pid:768217275 eyr:2020 byr:1937 hcl:#866857 hgt:157cm", "",
				"cid:270 byr:1993 hcl:#733820 ecl:hzl pid:722650020 hgt:174cm iyr:2010 eyr:2021", "",
				"hcl:#c0946f ecl:blu hgt:154cm eyr:2022 byr:1929 pid:357023679 iyr:2010", "",
				"ecl:hzl iyr:2013 hgt:165cm byr:1979 eyr:2023 hcl:#733820 pid:008734536", "",
				"hcl:#341e13 eyr:2030 byr:1993 iyr:2014 hgt:193cm cid:346 ecl:blu pid:536339538", "",
				"eyr:2030 ecl:hzl cid:296 pid:660062554 hcl:#efcc98 byr:1977 hgt:179cm iyr:2010", "",
				"eyr:2025 iyr:2010 hcl:#efcc98 pid:196372989 hgt:181cm byr:1952 ecl:oth", "",
				"cid:317 eyr:2026 ecl:blu hcl:#733820 hgt:184cm pid:549730813 byr:1927 iyr:2018", "",
				"pid:591769824 hgt:180cm byr:1920 ecl:blu eyr:2021 hcl:#cfa07d iyr:2017", "",
				"pid:988946348 hgt:183cm cid:117 byr:1955 ecl:blu iyr:2015 hcl:#623a2f eyr:2029", "",
				"byr:1995 hgt:182cm ecl:brn hcl:#6b5442 iyr:2012 eyr:2028 pid:482757872", "",
				"iyr:2017 cid:333 ecl:gry hcl:#623a2f hgt:157cm eyr:2021 pid:487895819 byr:1951", "",
				"hcl:#fffffd hgt:193cm eyr:2025 byr:1927 iyr:2014 ecl:oth pid:989206297", "",
				"eyr:2030 ecl:brn hcl:#18171d hgt:193cm iyr:2013 byr:1953 pid:862636088", "",
				"hcl:#fffffd pid:204286737 ecl:gry byr:1923 hgt:181cm iyr:2015 eyr:2023", "",
				"cid:288 pid:413935643 ecl:gry iyr:2012 hgt:171cm hcl:#623a2f eyr:2020 byr:1943", "",
				"hcl:#18171d ecl:gry pid:401957684 eyr:2020 iyr:2017 cid:141 byr:1944 hgt:74in", "",
				"pid:727198487 hgt:173cm cid:323 hcl:#18171d iyr:2012 eyr:2024 byr:1995 ecl:blu", "",
				"ecl:amb hcl:#602927 pid:460274414 hgt:76in byr:1995 iyr:2020 eyr:2028", "",
				"byr:2002 ecl:oth pid:101164770 hgt:172cm hcl:#fffffd eyr:2023 iyr:2016", "",
				"ecl:blu hcl:#888785 iyr:2016 pid:031162631 eyr:2025 hgt:186cm byr:1959", "",
				"ecl:blu pid:093242619 hgt:188cm byr:1970 eyr:2025 hcl:#6b5442 iyr:2020", "",
				"byr:1990 eyr:2025 ecl:grn pid:907309460 iyr:2011 hcl:#602927 hgt:62in", "",
				"pid:346468647 eyr:2021 ecl:oth hgt:169cm iyr:2010 cid:233 hcl:#b6652a byr:1977", "",
				"pid:904834317 iyr:2011 hcl:#b6652a eyr:2028 cid:281 byr:1944 hgt:187cm ecl:gry", "",
				"hgt:185cm eyr:2022 hcl:#341e13 ecl:oth byr:1934 pid:863541754 cid:178 iyr:2016", "",
				"eyr:2029 iyr:2014 byr:1937 cid:232 hgt:177cm hcl:#fffffd ecl:blu pid:076753558", "",
				"hcl:#cfa07d hgt:168cm ecl:grn pid:664159349 eyr:2028 iyr:2017 byr:1972", "",
				"hcl:#a97842 byr:1987 eyr:2020 hgt:182cm iyr:2018 ecl:brn pid:560272731", "",
				"pid:951007276 cid:260 eyr:2025 ecl:brn iyr:2015 byr:1957 hcl:#4b8216 hgt:161cm", "",
				"pid:359973697 hcl:#6b5442 eyr:2022 hgt:169cm byr:1965 ecl:brn iyr:2013", "",
				"iyr:2012 hgt:65in eyr:2024 pid:842371195 ecl:amb hcl:#341e13 byr:2000", "",
				"hgt:159cm ecl:amb hcl:#c0946f eyr:2020 pid:010539976 iyr:2011 byr:1921", "",
				"hgt:176cm cid:270 pid:838338992 eyr:2024 hcl:#866857 ecl:amb iyr:2015 byr:1982", "",
				"ecl:blu cid:246 hgt:185cm byr:1987 hcl:#fffffd pid:042361456 eyr:2022 iyr:2010", "",
				"hgt:164cm pid:881486702 ecl:brn byr:1969 hcl:#c0946f iyr:2010 eyr:2030", "",
				"iyr:2019 hcl:#6b5442 hgt:167cm ecl:amb cid:207 byr:1922 eyr:2025 pid:343956182", "",
				"byr:1988 pid:030965463 hgt:154cm ecl:gry eyr:2020 cid:227 iyr:2012 hcl:#3edc53", "",
				"hgt:158cm pid:270264980 eyr:2027 iyr:2016 byr:1928 cid:259 ecl:gry hcl:#733820", "",
				"hcl:#733820 cid:244 ecl:gry iyr:2013 eyr:2028 pid:794178180 hgt:74in byr:1923", "",
				"hcl:#a97842 byr:1934 ecl:hzl eyr:2027 pid:401882857 iyr:2018 hgt:185cm", "",
				"iyr:2018 pid:665564950 byr:1990 ecl:hzl hgt:154cm eyr:2026 hcl:#623a2f", "",
				"ecl:grn pid:734161280 hgt:184cm iyr:2018 eyr:2020 byr:1929 hcl:#a97842", "",
				"iyr:2018 byr:1925 eyr:2022 hgt:193cm ecl:hzl hcl:#341e13 pid:008582320", "",
				"hcl:#cfa07d ecl:hzl eyr:2029 cid:194 byr:1936 iyr:2020 hgt:186cm pid:328573727", "",
				"iyr:2011 hgt:188cm pid:338435675 cid:326 ecl:gry eyr:2027 hcl:#6b5442 byr:1958", "",
				"eyr:2020 hcl:#18171d byr:1978 iyr:2012 hgt:68in ecl:amb cid:346 pid:332495922", "",
				"ecl:blu hgt:61in pid:747650669 byr:1961 eyr:2028 iyr:2020 hcl:#4992f2", "",
				"byr:1958 iyr:2017 ecl:oth hgt:153cm hcl:#602927 eyr:2023 pid:108391213", "",
				"byr:1976 eyr:2023 iyr:2015 hgt:177cm pid:391628371 hcl:#8069c4 ecl:grn", "",
				"byr:1978 pid:302223240 iyr:2017 hgt:174cm hcl:#6b6569 ecl:blu eyr:2027", "",
				"cid:135 byr:1995 iyr:2015 ecl:oth pid:054611703 eyr:2023 hcl:#7d3b0c hgt:75in", "",
				"byr:1946 hgt:70in eyr:2022 hcl:#6b5442 ecl:amb iyr:2018 pid:859762925", "",
				"eyr:2028 hgt:162cm byr:1989 hcl:#0bd11f iyr:2020 ecl:gry pid:073498924", "",
				"iyr:2014 pid:122787281 byr:1982 cid:138 eyr:2021 hcl:#866857 ecl:hzl hgt:184cm", "",
				"eyr:2021 hgt:170cm cid:74 iyr:2019 pid:943445928 byr:1980 ecl:oth hcl:#ceb3a1", "",
				"iyr:2020 eyr:2030 pid:201122734 cid:246 hgt:169cm ecl:grn hcl:#fffffd byr:1962", "",
				"pid:025560194 byr:1989 hcl:#cfa07d hgt:182cm ecl:blu eyr:2025 iyr:2012", "",
				"hgt:151cm hcl:#efcc98 ecl:blu byr:1983 eyr:2023 pid:814513328 iyr:2013 cid:73", "",
				"byr:1961 pid:536384108 hgt:188cm ecl:amb iyr:2013 eyr:2027 hcl:#888785 cid:121", "",
				"pid:364607819 eyr:2024 ecl:amb hcl:#b6652a iyr:2016 byr:2000 hgt:187cm", "",
				"cid:207 pid:913509058 ecl:brn byr:2001 eyr:2026 hcl:#866857 iyr:2019 hgt:180cm", "",
				"pid:363979129 eyr:2027 iyr:2013 ecl:gry hcl:#866857 byr:1957 hgt:62in", "",
				"byr:1932 eyr:2027 hgt:66in ecl:hzl hcl:#efcc98 pid:417620217 iyr:2013", "",
				"byr:1960 hcl:#888785 hgt:176cm ecl:hzl pid:025206542 iyr:2015 eyr:2030", "",
				"ecl:oth hgt:182cm hcl:#341e13 pid:526568190 iyr:2018 cid:280 byr:1997 eyr:2028", "",
				"hgt:186cm pid:273625601 byr:1993 iyr:2018 eyr:2021 hcl:#733820 ecl:blu", "",
				"hcl:#a97842 pid:347084450 ecl:oth eyr:2030 hgt:176cm byr:1955 cid:229 iyr:2013", "",
				"hcl:#fffffd byr:1979 iyr:2017 pid:183840860 hgt:177cm ecl:blu eyr:2023", "",
				"pid:273620987 eyr:2022 hgt:162cm cid:269 byr:1991 hcl:#602927 ecl:amb iyr:2019", "",
				"eyr:2022 byr:1988 hgt:190cm pid:349839553 hcl:#602927 iyr:2018 ecl:gry", "",
				"iyr:2014 ecl:gry hcl:#733820 eyr:2025 hgt:179cm pid:231854667 byr:1984 cid:102", "",
				"eyr:2020 pid:509400891 hcl:#cfa07d hgt:172cm ecl:grn byr:1997 iyr:2020", "",
				"iyr:2017 byr:1994 hgt:174cm ecl:amb pid:685743124 hcl:#fffffd eyr:2029", "",
				"iyr:2012 hgt:177cm byr:1999 pid:549190825 hcl:#b6652a eyr:2028 ecl:oth cid:316", "",
				"hgt:192cm ecl:grn byr:1924 iyr:2011 eyr:2029 hcl:#efcc98 pid:215962187", "",
				"iyr:2011 hcl:#866857 cid:164 hgt:184cm ecl:gry eyr:2023 byr:1959 pid:204093118", "",
				"hgt:172cm ecl:hzl hcl:#3f2f3a pid:623470811 byr:1938 iyr:2013 eyr:2022", "",
				"ecl:oth hcl:#602927 pid:049746898 byr:1924 hgt:150cm eyr:2026 iyr:2014", "",
				"eyr:2023 cid:308 hgt:170cm ecl:oth iyr:2014 hcl:#18171d pid:874405208 byr:1936", "",
				"eyr:2021 ecl:hzl pid:423603306 hcl:#c0946f cid:147 byr:1988 iyr:2016 hgt:164cm", "",
				"hgt:176cm iyr:2010 hcl:#6b5442 cid:280 byr:1988 ecl:hzl pid:967151288 eyr:2028", "",
				"cid:299 hgt:163cm ecl:gry pid:561439154 eyr:2023 hcl:#cfa07d iyr:2019 byr:1959", "",
				"pid:635547007 ecl:blu byr:1996 hcl:#7d3b0c cid:280 eyr:2023 hgt:170cm iyr:2017", "",
				"hcl:#c0946f cid:199 hgt:162cm ecl:amb pid:130696599 eyr:2022 iyr:2018 byr:1948", "",
				"cid:314 hcl:#a4fc09 ecl:hzl iyr:2019 pid:886849824 eyr:2026 byr:1933 hgt:178cm", "",
				"byr:1996 iyr:2016 eyr:2030 hgt:169cm pid:119207760 hcl:#ef542c ecl:brn", "",
				"pid:727812879 iyr:2013 eyr:2027 hgt:172cm hcl:#7d3b0c ecl:gry byr:1966", "",
				"hcl:#341e13 iyr:2016 pid:744997238 cid:322 byr:1973 ecl:hzl eyr:2028 hgt:190cm", "",
				"hgt:171cm eyr:2026 iyr:2014 ecl:oth pid:074049558 hcl:#04083f byr:1923", "",
				"pid:973713235 eyr:2021 ecl:brn byr:1922 hcl:#fffffd iyr:2012 hgt:178cm", "",
				"ecl:blu hgt:165cm iyr:2012 eyr:2025 pid:775787557 byr:1952 hcl:#623a2f", "",
				"iyr:2010 eyr:2029 hgt:74in ecl:hzl byr:1926 pid:348573885 hcl:#9d1214", "",
				"hgt:171cm ecl:oth eyr:2022 pid:148728436 byr:1993 hcl:#a97842 iyr:2013", "",
				"iyr:2019 hgt:151cm eyr:2020 pid:319882814 ecl:grn byr:1966 cid:256 hcl:#3107b3", "",
				"hgt:184cm ecl:grn byr:1947 eyr:2025 iyr:2015 pid:827962962 cid:62 hcl:#f3a364", "",
				"iyr:2013 hcl:#fffffd pid:215012801 ecl:amb eyr:2024 hgt:154cm byr:1973", "",
				"ecl:hzl hgt:152cm hcl:#623a2f byr:1944 eyr:2022 pid:295632731 cid:243 iyr:2019", "",
				"ecl:gry pid:445116331 hcl:#a97842 hgt:187cm eyr:2026 iyr:2020 byr:1992", "",
				"pid:058987865 byr:1927 cid:209 hcl:#65ccf6 eyr:2025 ecl:brn iyr:2012 hgt:164cm", "",
				"eyr:2028 iyr:2013 pid:103268377 hgt:179cm byr:1922 ecl:hzl hcl:#7d3b0c", "",
				"byr:1923 ecl:gry hgt:167cm hcl:#7fc8ee iyr:2015 pid:427963077 eyr:2024", "",
				"byr:1927 ecl:grn pid:741328150 eyr:2029 hcl:#733820 iyr:2015 hgt:157cm", "",
				"hcl:#623a2f eyr:2020 byr:1936 ecl:gry pid:236984204 iyr:2011 hgt:156cm", "",
				"pid:774489073 iyr:2013 byr:1922 ecl:brn eyr:2025 hcl:#18171d hgt:163cm", "",
				"eyr:2024 hgt:65in byr:1962 iyr:2019 pid:112233558 hcl:#888785 ecl:grn", "",
				"hgt:172cm eyr:2022 hcl:#18171d ecl:blu pid:609008608 iyr:2013 cid:244 byr:1980", "",
				"byr:1926 hgt:188cm ecl:hzl eyr:2021 iyr:2018 hcl:#866857 pid:557800355", "",
				"byr:1939 pid:200409089 eyr:2026 hgt:164cm ecl:grn iyr:2013 hcl:#733820", "",
				"pid:609818712 hcl:#733820 byr:1958 eyr:2025 hgt:187cm iyr:2017 ecl:gry", "",
				"hgt:66in pid:618590610 iyr:2013 byr:1938 hcl:#d1bda9 eyr:2022 ecl:grn cid:69", "",
				"hgt:156cm pid:755742405 byr:1929 hcl:#6b5442 eyr:2024 iyr:2018 ecl:gry cid:105", "",
				"iyr:2015 hcl:#341e13 byr:1968 pid:434159843 ecl:amb hgt:150cm eyr:2030", "",
				"cid:249 hcl:#cfa07d hgt:180cm ecl:gry eyr:2026 byr:1965 pid:048327438 iyr:2010", "",
				"pid:136468890 ecl:gry byr:1940 hcl:#fffffd hgt:185cm iyr:2016 eyr:2021", "",
				"iyr:2010 byr:1932 eyr:2025 ecl:grn pid:595837820 hcl:#341e13 hgt:166cm cid:224", "",
				"pid:481646831 eyr:2029 hcl:#623a2f cid:319 iyr:2016 ecl:brn hgt:160cm byr:1944", "",
				"iyr:2014 hgt:170cm byr:1963 hcl:#623a2f eyr:2026 pid:225910806 ecl:gry", "",
				"ecl:grn iyr:2010 hgt:193cm byr:1928 eyr:2028 pid:343022641 hcl:#733820", "",
				"eyr:2023 ecl:grn byr:1950 iyr:2012 hcl:#866857 pid:400725165 hgt:193cm", "",
				"cid:195 iyr:2014 ecl:oth eyr:2027 byr:1966 hgt:177cm hcl:#18171d pid:913894485", "",
				"iyr:2015 hgt:154cm cid:206 pid:134599284 hcl:#602927 eyr:2023 ecl:brn byr:1983", "",
				"hgt:189cm byr:1965 iyr:2011 cid:178 ecl:hzl hcl:#b6652a eyr:2026 pid:683560227", "",
				"eyr:2030 pid:047446524 ecl:grn hgt:167cm iyr:2017 hcl:#602927 byr:1920", "",
				"byr:1991 ecl:grn iyr:2016 hcl:#5e1ef2 hgt:186cm pid:076499738 eyr:2025", "",
				"eyr:2030 hcl:#18171d pid:750694893 hgt:157cm iyr:2020 cid:338 byr:1956 ecl:gry", "",
				"hgt:163cm byr:1962 cid:108 ecl:gry hcl:#733820 iyr:2012 eyr:2029 pid:763684725", "",
				"byr:1984 hcl:#888785 hgt:159cm iyr:2012 ecl:gry eyr:2024 cid:236 pid:174711749", "",
				"byr:1973 iyr:2018 hcl:#a97842 pid:937214113 ecl:blu cid:247 hgt:186cm eyr:2023", "",
				"pid:298274796 byr:1928 hcl:#a97842 hgt:188cm iyr:2011 ecl:gry eyr:2028", "",
				"pid:499116613 eyr:2024 ecl:gry hcl:#cfa07d hgt:193cm byr:1999 cid:278 iyr:2015", "",
				"hcl:#6b5442 eyr:2027 hgt:175cm byr:1988 ecl:brn pid:410075320 iyr:2010", "",
				"ecl:oth pid:144593265 hcl:#fffffd eyr:2020 iyr:2018 byr:1975 hgt:160cm cid:304", "",
				"iyr:2014 hcl:#ceb3a1 eyr:2029 byr:1951 pid:520804395 hgt:185cm ecl:oth", "",
				"hgt:159cm pid:312887994 cid:205 iyr:2016 ecl:hzl hcl:#866857 eyr:2029 byr:1944", "",
				"byr:1946 eyr:2023 hgt:163cm hcl:#2d4e9c ecl:brn pid:839043333 iyr:2014", "",
				"eyr:2029 hgt:189cm ecl:blu byr:1922 iyr:2016 pid:924104599 hcl:#b6652a", "",
				"cid:118 pid:959515596 byr:1921 iyr:2010 eyr:2029 hcl:#7d3b0c ecl:oth hgt:158cm", "",
				"iyr:2013 byr:1935 hgt:179cm eyr:2023 ecl:amb pid:073621563 hcl:#623a2f", "",
				"hcl:#18171d cid:230 byr:1989 ecl:oth eyr:2021 hgt:181cm pid:661224730 iyr:2019", "",
				"ecl:hzl pid:758144605 hcl:#ceb3a1 hgt:186cm eyr:2028 iyr:2014 byr:1928", "",
				"ecl:hzl hgt:66in byr:2000 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", "",
				"iyr:2012 pid:749770535 byr:1969 cid:148 hcl:#733820 hgt:180cm eyr:2021 ecl:hzl", "",
				"iyr:2010 byr:1958 hgt:164cm ecl:blu hcl:#733820 pid:890634327 eyr:2024", "",
				"hgt:70in pid:218397894 iyr:2020 eyr:2025 ecl:gry hcl:#341e13 byr:1970", "",
				"eyr:2020 pid:854208004 hgt:157cm hcl:#7d3b0c ecl:amb byr:1981 iyr:2020", "",
				"byr:1924 cid:321 eyr:2028 hcl:#cfa07d iyr:2010 ecl:amb pid:036669613 hgt:170cm", "",
				"hcl:#ceb3a1 ecl:grn eyr:2028 pid:074363489 iyr:2010 hgt:173cm byr:1966", "",
				"hgt:175cm iyr:2016 eyr:2024 cid:244 byr:1952 pid:085432899 hcl:#fffffd ecl:brn", "",
				"ecl:brn eyr:2026 iyr:2017 hgt:75in pid:745302991 byr:1969 hcl:#7394c7", "",
			},
			expected1: 187,
			expected2: 186,
		},
	}

	for i, tt := range tests {
		fmt.Printf("\nCase %d\n", i)
		result1, result2, err := day4(tt.input)
		if err != nil {
			t.Errorf("error in case %d: %v", i, err)
		}
		if result1 != tt.expected1 {
			t.Errorf("error solving case %d part 1, got: %d, want: %d", i, result1, tt.expected1)
		}
		if result2 != tt.expected2 {
			t.Errorf("error solving case %d part 2, got: %d, want: %d", i, result2, tt.expected2)
		}
	}
}
