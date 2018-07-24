package ballparser

import "testing"

func TestBallToString(t *testing.T) {
	tests := []struct{in, out string}{
		{
			in:`<tr week="4" style="background: none;"><td>2018074</td><td>2018-06-28(四)</td><td><span class="ball_5">09</span>&nbsp;<span class="ball_5">11</span>&nbsp;<span class="ball_5">14</span>&nbsp;<span class="ball_5">20</span>&nbsp;<span class="ball_5">27</span>&nbsp;<span class="ball_5">30</span>&nbsp;</td><td><span class="ball_1">09<span></span></span></td><td>327,800,026</td><td>9</td><td>6,057,912</td><td>120</td><td>99,179</td><td>8亿1358万1190元</td><td><a target="_blank" href="http://cp.360.cn/experience/ssq?Issue=2018074">查看统计</a></td></tr>`,
			out:"091114202730|09",
		},
	}

	for _,tt:=range tests{
		if actual, err := BallToString(tt.in); tt.out != actual || err !=nil{
			t.Errorf("get %s, excepted %s", actual, tt.out)
		}
	}
}
