package example

import "testing"

// cd example
// go test
// go test -v 会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率
// go test -run TestAdd -v 只想运行其中的一个用例
func TestAdd(t *testing.T) {
	if ans := Add(1, 3); ans != 4 {
		t.Errorf("1 + 3 expected be 4, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != 6 {
			t.Fatal("fail")
		}
	})
}

type s struct {
	name           string
	a, b, expacted int
}

func createTests(t *testing.T, s *s) {
	t.Helper() //让报错信息更准确，有助于定位
	if ans := Mod(s.a, s.b); ans != s.expacted {
		t.Fatalf("%d / %d expected %d, but %d got",
			s.a, s.b, s.expacted, ans)
	}
}

func TestMod(t *testing.T) {
	createTests(t, &s{name:"1", a: 2, b: 1, expacted: 2})
	createTests(t, &s{name:"2", a: 2, b: 2, expacted: 1})
	createTests(t, &s{name:"3", a: 3, b: 2, expacted: 0})
	createTests(t, &s{name:"4", a: 3, b: 2, expacted: 1})
}
