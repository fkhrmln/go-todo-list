package helper

import "go-todo-list/dto/request"

func CheckValueUpdateTodoRequest(updateTodoRequest request.UpdateTodoRequest) map[string]interface{} {
	updateTodoRequestMap := map[string]interface{}{}

	if updateTodoRequest.Title != "" {
		updateTodoRequestMap["title"] = updateTodoRequest.Title
	}

	if updateTodoRequest.Body != "" {
		updateTodoRequestMap["body"] = updateTodoRequest.Body
	}

	if updateTodoRequest.Deadline != 0 {
		updateTodoRequestMap["deadline"] = updateTodoRequest.Deadline
	}

	if updateTodoRequest.IsCompleted {
		updateTodoRequestMap["is_completed"] = updateTodoRequest.IsCompleted
	}

	return updateTodoRequestMap
}
