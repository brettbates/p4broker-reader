package reader

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// Read dictionary pairs from the output of a p4broker filter command
// Stopping when all args are read in
// Pass os.Stdin to read p4brokers output directly
func Read(stdin io.Reader) (map[string]string, error) {
	out := make(map[string]string)
	reader := bufio.NewReader(stdin)
	var end string
	for {
		line, err := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if err != nil {
			log.Printf("Failed to read from stdin, err:\n\t%s", err)
			return nil, err
		}
		// TODO check we have a x: y line
		split := strings.Split(line, ": ")
		if split[0] == "argCount" {
			cnt, err := strconv.Atoi(split[1])
			cnt--
			if err != nil {
				log.Printf("Failed to parse argCount %s: %s", split[0], split[1])
				return nil, err
			}
			end = "Arg" + fmt.Sprintf("%d", cnt)
		}

		out[split[0]] = split[1]

		if len(end) > 0 && split[0] == end {
			break
		}
	}
	return out, nil
}
