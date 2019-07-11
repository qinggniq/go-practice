package main

import {
	"os"
	"bufio"
}
func countLines(f *os.File, counts map[string]int) {
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		counts[scan.Text()]++;
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if (len(args) == 0) {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := counts {
		fmt.Printf("%d, %s\n", line, n);
	}
	
}