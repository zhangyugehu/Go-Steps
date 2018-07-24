package ballparser

import (
	"regexp"
)

// var re360 = regexp.MustCompile(`<td>([0-9]{7})</td>.*<td><span class="ball_5">([0-9]{2})</span>.*<span class="ball_5">([0-9]{2})</span>.*<span class="ball_5">([0-9]{2})</span>.*<span class="ball_5">([0-9]{2})</span>.*<span class="ball_5">([0-9]{2})</span>.*<span class="ball_5">([0-9]{2})</span>.*</td><td><span class="ball_1">([0-9]{2})<span>`)
var re360 = regexp.MustCompile(`<tr week='[0-9]'><td>([0-9]{7})</td><td>[0-9]{4}-[0-9]{2}-[0-9]{2}\([二四日]\)</td><td><span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;</td><td><span class='ball_1'>([0-9]{2})<span></td><td>[0-9,\- ]+</td><td>[0-9,\- ]+</td><td>[0-9,\- ]+</td><td>[0-9\- ]+</td><td>[0-9,\- ]+</td><td>[0-9亿万千百十元\- ]+</td><td><a target='_blank' href='http://cp\.360\.cn/experience/ssq\?Issue=[0-9]{7}'>查看统计</a></td></tr>`)

func BallToString(text string) (map[string]int, error) {

	allMatches := re360.FindAllStringSubmatch(text, -1)

	codeMap := make(map[string]int)

	for i, _ := range allMatches{
		//for j, n:=range m{
		//	if j==0 {
		//		continue
		//	}
		//	fmt.Println(j, n)
		//}
		code :=
			//allMatches[i][0]+ " " +
			allMatches[i][2]+
			allMatches[i][3]+
			allMatches[i][4]+
			allMatches[i][5]+
			allMatches[i][6]+
			allMatches[i][7]+
			"|" +
			allMatches[i][8]

		if _,ok := codeMap[code]; ok{
			codeMap[code] ++
		}else{
			codeMap[code] = 1
		}
	}
	//fmt.Println(allMatches[0][2])
	//fmt.Println(allMatches[0][3])
	//fmt.Println(allMatches[0][4])
	//fmt.Println(allMatches[0][5])
	//fmt.Println(allMatches[0][6])
	//fmt.Println(allMatches[0][7])
	//fmt.Println(allMatches[0][8])
	return codeMap, nil
}