package eval

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Errorf("error parsing expression '%s': %s", test.expr, err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestString(t *testing.T) {
	// 测试用例列表
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
		"-1 + -x",
		"-x * -y",
	}

	for _, originalStr := range tests {
		fmt.Printf("\nOriginal: %s\n", originalStr)

		// 1. 第一次解析
		expr1, err := Parse(originalStr)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		// 2. 转换为字符串 (Pretty Print)
		generatedStr := expr1.String()
		fmt.Printf("String(): %s\n", generatedStr)

		// 3. 第二次解析 (解析生成的字符串)
		expr2, err := Parse(generatedStr)
		if err != nil {
			fmt.Println("Re-parse error:", err)
			continue
		}

		// 4. 比较两个表达式树是否结构一致
		// 注意：reflect.DeepEqual 会深度比较两个结构体
		if !reflect.DeepEqual(expr1, expr2) {
			fmt.Println("FAIL: 结构不一致!")
			// 调试用：打印两个对象的具体类型和内容
			fmt.Printf("Expr1: %#v\n", expr1)
			fmt.Printf("Expr2: %#v\n", expr2)
		} else {
			fmt.Println("PASS: 语法树一致")
		}
	}
}
