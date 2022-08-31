package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

type SearchCommand struct {
	Search string
}

type UpdateResponse struct {
	Update []ListItem
}

type ListItem struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Icon         Icon   `json:"icon"`
	CategoryIcon Icon   `json:"category_icon"`
}

type Icon struct {
	Name string
}

const popLauncher string = "/Users/a.lacoin/Developer/pomdtr/raycast-linux/bin/pop-launcher"

func launcher(searches chan SearchCommand, updates chan UpdateResponse) error {
	cmd := exec.Command(popLauncher)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdout.Close()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			var update UpdateResponse
			err := json.Unmarshal(scanner.Bytes(), &update)
			if err != nil {
				log.Println(err)
				continue
			}
			updates <- update
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stderr.Close()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		encoder := json.NewEncoder(stdin)
		for search := range searches {
			fmt.Println("search:", search.Search)
			encoder.Encode(search)
		}

		stdin.Close()
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}

	}()

	if err = cmd.Start(); err != nil {
		return err
	}

	return nil
}

func main() {
	searches := make(chan SearchCommand)
	updates := make(chan UpdateResponse)
	err := launcher(searches, updates)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		writer := io.Writer(os.Stdout)
		encoder := json.NewEncoder(writer)
		for update := range updates {
			encoder.Encode(update)
		}
	}()

	done := make(chan bool)
	go func() {
		searches <- SearchCommand{Search: "s"}
		time.Sleep(2 * time.Second)
		searches <- SearchCommand{Search: "t"}
		time.Sleep(2 * time.Second)
		searches <- SearchCommand{Search: "r"}
		time.Sleep(2 * time.Second)
		done <- true
	}()

	// sigs := make(chan os.Signal)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// select {
	// case <-done:
	// 	fmt.Println("finished")
	// case <-sigs:
	// 	fmt.Println("signal")
	// }

	<-done
	fmt.Println("exiting")
}
