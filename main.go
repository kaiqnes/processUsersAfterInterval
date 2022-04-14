package main

import (
	"fmt"
	data_structures "processUsersAfterInterval/data-structures"
	"processUsersAfterInterval/model"
)

const (
	minimumInterval = 4
	printDebugLogs  = false
)

var (
	userStream = model.User{7, 3, 10, 5, 10, 8, 3, 3, 3, 3, 9, 6, 2, 10}
)

func main() {
	/*
		dada uma entrada de usuários através de um stream
		[]int{7, 3, 10, 5, 10, 8, 9, 6, 2, 10}
		e um intervalo minimo J = 4
		imprima todas as ocorrencias de reprocessamento do mesmo usuário enquanto menor que intervalo minimo
	*/

	fmt.Println("Reprocess before minimum interval:", processUserStream(userStream))
}

func processUserStream(users model.User) model.User {
	queue := data_structures.NewQueue()
	processedUsers := map[int]int{}
	blockedUsers := model.User{}

	for index, user := range users {
		if printDebugLogs {
			fmt.Println("Iteration", index+1)
			fmt.Println("QueueSize:", queue.Size(), "| Queue:", queue)
			fmt.Println("User:", user)
			fmt.Printf("%+v\n\n", processedUsers)
		}
		if queue.Size() == minimumInterval {
			poppedUser := queue.Dequeue().(int)
			processedUsers[poppedUser]--

			if processedUsers[poppedUser] == 0 {
				delete(processedUsers, poppedUser)
			}
		}

		queue.Enqueue(user)
		isProcessed := processedUsers[user]
		if isProcessed > 0 {
			blockedUsers = append(blockedUsers, user)
		}
		processedUsers[user]++
	}

	return blockedUsers
}
