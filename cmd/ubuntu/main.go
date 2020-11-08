package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/igor-gerasimovich/mg-track-time-pc/utils"
)

func main() {
	ctx := context.Background()

	for {
		windowID, err := getWindowID(ctx)
		if err != nil {
			log.Fatal(err)
			continue
		}

		name, classes, err := getWindowNameClass(ctx, windowID)
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Print("name: ", string(name), " | ")
		fmt.Print("class: ")
		for _, c := range classes {
			fmt.Print(string(c), " ")
		}

		fmt.Print("\n")

		time.Sleep(1 * time.Second)
	}
}

func getWindowID(ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "xprop", "-root", "_NET_ACTIVE_WINDOW")
	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	windowID := utils.SplitLastByte(res, '#')
	return utils.FilterByte(windowID, ' '), nil
}

func getWindowNameClass(ctx context.Context, windowID []byte) ([]byte, [][]byte, error) {
	cmd := exec.CommandContext(ctx, "xprop", "-id", string(windowID), "_NET_WM_NAME", "WM_CLASS")
	res, err := cmd.Output()
	if err != nil {
		return nil, nil, err
	}

	res = bytes.Trim(res, "\n")
	parts := bytes.Split(res, []byte("\n"))
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("returned not valid count of data. must have 2, got %d", len(parts))
	}

	wmName := utils.BytesCoveredRune(parts[0], '"')
	if len(wmName) != 1 {
		return nil, nil, fmt.Errorf("failed to get wmname must have 1 value, got %d", len(wmName))
	}

	wmClass := utils.BytesCoveredRune(parts[1], '"')
	if len(wmName) != 1 {
		return nil, nil, fmt.Errorf("failed to get wmname must have 2 value, got %d", len(wmName))
	}

	return wmName[0], wmClass, nil
}
