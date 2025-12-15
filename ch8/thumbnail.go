package main

import (
	"log"
	"os"
	"sync"
)

func ImageFile(infile string) (string, error) {
	return infile, nil
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, filename := range filenames {
		go ImageFile(filename)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	// 闭包陷阱
	for _, filename := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	ch := make(chan error)
	for _, filename := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			ch <- err
		}(filename)
	}

	for range filenames {
		// 这里有一个问题： 由于channel是无缓存的，第一次出现错误之后直接就会返回err，后续出现错误就会导致对应的goroutine阻塞
		if err := <-ch; err != nil {
			return err
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	// 通过“保证每个 Goroutine 都有写入空间”，成功避免了因发送阻塞而导致的 Goroutine 僵死泄露
	ch := make(chan item, len(filenames))
	for _, filename := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(filename)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for filename := range filenames {
		wg.Add(1) // 这里要在main goroutine中完成处理，因为如果是在内部，可能go还没有被调度，就执行到了wait(),直接结束
		go func(f string) {
			defer wg.Done() // 确保 add（-1）
			thumbfile, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()
		}(filename)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var sum int64
	for size := range sizes {
		sum += size
	}
	return sum
}
