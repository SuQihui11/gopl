package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

// 简单的清屏函数，为了让显示看起来像个真正的“墙”
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Clock struct {
	Name    string
	Address string
	Time    string
	Failed  bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clockwall NAME=HOST:PORT ...")
		return
	}
	// 初始化
	var clocks []*Clock
	for _, arg := range os.Args[1:] {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			continue
		}
		clocks = append(clocks, &Clock{
			Name:    parts[0],
			Address: parts[1],
			Time:    "Connecting......",
		})
	}

	// 设置锁，确保并发读写
	var mux sync.Mutex

	// 建立连接获取时间显示
	for _, clock := range clocks {
		go func(clock *Clock) {
			// 建立链接
			conn, err := net.Dial("tcp", clock.Address)
			if err != nil {
				mux.Lock()
				clock.Failed = true
				clock.Time = "Connection Failed"
				mux.Unlock()
			}
			defer conn.Close()
			// 通过scanner从conn读取数据
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				mux.Lock()
				clock.Time = scanner.Text()
				mux.Unlock()
			}

			mux.Lock()
			clock.Failed = true
			clock.Time = "Disconnected"
			mux.Unlock()
		}(clock)
	}

	// 设置ticker计时器
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// 初始化 TabWriter 用于对齐表格
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	for range ticker.C {
		// 清屏
		clearScreen()

		// 打印表头
		fmt.Fprintln(w, "LOCATION\tTIME\tSTATUS")
		fmt.Fprintln(w, "--------\t----\t------")

		mux.Lock()
		// 按照名字排序，保证显示顺序固定（可选）
		sort.Slice(clocks, func(i, j int) bool {
			return clocks[i].Name < clocks[j].Name
		})

		for _, c := range clocks {
			status := "OK"
			if c.Failed {
				status = "ERROR"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", c.Name, c.Time, status)
		}
		mux.Unlock()

		w.Flush() // 将缓冲的内容输出到屏幕
	}
}
