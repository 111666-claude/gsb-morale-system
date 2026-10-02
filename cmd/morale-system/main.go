// Command morale-system 跑士气样例。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"example.com/morale"
)

// Run 执行一次命令行调用，返回退出码。
func Run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("morale-system", flag.ContinueOnError)
	flags.SetOutput(stderr)
	sample := flags.String("sample", "break", "break / rally / recover / kill / scan")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	switch *sample {
	case "break":
		field := morale.New(50, 20, 30, 5, 60000)
		field.Add("u1", 0, 0, false)
		field.Damage("u1", 800, 0)
		fmt.Fprintf(stdout, "broken=%v morale=%d\n", field.Broken("u1"), field.Morale("u1"))
	case "rally":
		field := morale.New(50, 20, 30, 5, 60000)
		field.Add("u1", 0, 0, false)
		field.Damage("u1", 800, 0)
		field.Heal("u1", 9)
		fmt.Fprintf(stdout, "broken=%v morale=%d\n", field.Broken("u1"), field.Morale("u1"))
	case "recover":
		field := morale.New(50, 20, 30, 5, 60000)
		field.Add("u1", 0, 0, false)
		field.Damage("u1", 100, 0)
		for index := 1; index <= 10; index++ {
			field.Tick(int64(index) * 1000)
		}
		fmt.Fprintf(stdout, "morale=%d\n", field.Morale("u1"))
	case "kill":
		field := morale.New(50, 20, 30, 5, 60000)
		field.Add("a", 0, 0, false)
		field.Add("b", 1000, 0, false)
		field.Kill("a", 0)
		fmt.Fprintf(stdout, "b=%d\n", field.Morale("b"))
	case "scan":
		field := morale.New(50, 20, 30, 5, 60000)
		for index := 0; index < 3000; index++ {
			field.Add(fmt.Sprintf("u-%d", index), float64(index), 0, false)
		}
		for index := 0; index < 3000; index++ {
			field.Kill("u-2999", int64(index))
		}
		fmt.Fprintf(stdout, "scanned=%d\n", field.Scanned())
	default:
		fmt.Fprintln(stderr, "需要 --sample break|rally|recover|kill|scan")
		return 2
	}
	return 0
}

func main() {
	os.Exit(Run(os.Args[1:], os.Stdout, os.Stderr))
}
