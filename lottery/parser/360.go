package parser


import "regexp"


type Lottery360 struct {

}

func (p *Lottery360) Ssq(byte []byte) (ParseResult, error) {
	reSsq := regexp.MustCompile(`<tr week='[0-9]'><td>([0-9]{7})</td><td>[0-9]{4}-[0-9]{2}-[0-9]{2}\([二四日]\)</td><td><span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;<span class='ball_5'>([0-9]{2})</span>&nbsp;</td><td><span class='ball_1'>([0-9]{2})<span></td><td>[0-9,\- ]+</td><td>[0-9,\- ]+</td><td>[0-9,\- ]+</td><td>[0-9\- ]+</td><td>[0-9,\- ]+</td><td>[0-9亿万千百十元\- ]+</td><td><a target='_blank' href='http://cp\.360\.cn/experience/ssq\?Issue=[0-9]{7}'>查看统计</a></td></tr>`)
	allMatches := reSsq.FindAllSubmatch(byte, -1)
	codeMap := make(map[string]int)
	for _, match := range allMatches{
		code := string(match[2]) +
			string(match[3])+
			string(match[4])+
			string(match[5])+
			string(match[6])+
			string(match[7])+
			"|" +
			string(match[8])

		if _,ok := codeMap[code]; ok{
			codeMap[code] ++
		}else{
			codeMap[code] = 1
		}
	}
	return ParseResult{codeMap}, nil
}

func (p *Lottery360) Dlt(byte []byte) (ParseResult, error) {
	// TODO
	return ParseResult{}, nil
}

func (p *Lottery360) Fc3d(byte []byte) (ParseResult, error) {
	reFc3d := regexp.MustCompile(`<tr week='[0-9]'><td>([0-9]{7})</td><td>[0-9]{4}-[0-9]{2}-[0-9]{2}\([一二三四五六日]\)</td><td><span class='ball_5'>([0-9])</span>&nbsp;<span class='ball_5'>([0-9])</span>&nbsp;<span class='ball_5'>([0-9])</span>&nbsp;\(组[三六]\)<td>[0-9]{3}</td><td>[0-9]{2},[0-9]{3},[0-9]{3}</td><td>[0-9]{5}</td><td>[0-9],[0-9]{3}</td><td>[0-9]{1,5}</td><td>[0-9]{3}</td><td>[0-9]{1,5}</td><td>[0-9]{3}</td><td><a target='_blank' href='http://cp\.360\.cn/experience/sd\?Issue=[0-9]{7}'>查看统计</a></td></tr>`)
	allMatches := reFc3d.FindAllSubmatch(byte, -1)
	codeMap := make(map[string]int)
	for _, match := range allMatches{
		code := string(match[2]) +
			string(match[3])+
			string(match[4])

		if _,ok := codeMap[code]; ok{
			codeMap[code] ++
		}else{
			codeMap[code] = 1
		}
	}
	return ParseResult{codeMap}, nil
}

func (p *Lottery360) Q3(byte []byte) (ParseResult, error) {
	reQ3 := regexp.MustCompile(`<span class='sum'>(\d{1,2})</span></td><td class='tdbdr'></td><td class='tdbg_1' ><span class='sum'>(\d{1,2})</span>`)
	allMatches := reQ3.FindAllSubmatch(byte, -1)

	codeMap := make(map[string]int)
	for _, match := range allMatches{
		code := string(match[1])

		if _,ok := codeMap[code]; ok{
			codeMap[code] ++
		}else{
			codeMap[code] = 1
		}
	}
	return ParseResult{codeMap}, nil
}
