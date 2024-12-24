package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
*/

func MakeCommands(commands []string) {
	for _, command := range commands {
		commandSlice := strings.Fields(command)
		if len(commandSlice) == 0 {
			continue
		}
		if commandSlice[0] == "cd" {
			os.Chdir(commandSlice[1])
		} else if commandSlice[0] == "pwd" {
			path, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(path)
		} else if commandSlice[0] == "echo" {
			if len(commandSlice) <= 1 {
				fmt.Println()
			} else {
				fmt.Println(strings.Join(commandSlice[1:], " "))
			}
		} else if commandSlice[0] == "kill" {
			if len(commandSlice) <= 1 {
				log.Fatal("please input pid of process")
			}
			kill := exec.Command(commandSlice[0], commandSlice[1])
			kill.Stdin = os.Stdin
			kill.Stdout = os.Stdout
			err := kill.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Process with pid ", commandSlice[1], " killed succesfully")
			}
		} else {
			ps := exec.Command(commandSlice[0])
			ps.Stdout = os.Stdout
			ps.Stderr = os.Stderr
			err := ps.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " | ")
		MakeCommands(commands)
	}
}
