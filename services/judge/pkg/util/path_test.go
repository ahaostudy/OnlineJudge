package util

import (
	"fmt"
	"testing"
)

func TestRemoveDirectoryFromPath(t *testing.T) {
	fmt.Println(RemoveDirectoryFromPath("/dir/a/a.txt", "/dir/a"))
	fmt.Println(RemoveDirectoryFromPath("/dir/a/b/a.txt", "/dir/a"))
	fmt.Println(RemoveDirectoryFromPath("/dir/a", "/dir/a"))
	fmt.Println(RemoveDirectoryFromPath("/dir/a//a.txt", "/dir/a"))
	fmt.Println(RemoveDirectoryFromPath(`
File "/projects/OnlineJudge/data/code/78ea682d-643e-41c3-879d-d6e979e35e3e.py", line 1
    print(sum(map(int, input().split()))
         ^
SyntaxError: '(' was never closed
`, "/projects/OnlineJudge/data/code"))
}
