package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readList() ([]Task, error) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeList(tasks []Task) error {
	jsonTasks, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile("data.json", jsonTasks, 0644); err != nil {
		return err
	}
	return nil
}

func TaskAddFunc(task string) (string, error) {
	if task == "" {
		return "", errors.New("task cannot be empty")
	}

	tasks, err := readList()
	if err != nil {
		return "", fmt.Errorf("error with reading data.json: %v", err)
	}

	newTask := Task{
		ID:        int64(len(tasks) + 1),
		Desc:      task,
		Status:    "todo",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: "",
	}
	tasks = append(tasks, newTask)

	if err := writeList(tasks); err != nil {
		return "", fmt.Errorf("error with writing data.json: %v", err)
	}

	return fmt.Sprintf("Task will be added with ID: %d", newTask.ID), nil
}

func TaskUpdateFunc(taskID, desc string) error {
	tasks, err := readList()
	if err != nil {
		return fmt.Errorf("error with reading data.json: %v", err)
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("the id was expected to be a string: %v", err)
	}

	for idx := range tasks {
		if tasks[idx].ID == int64(id) {
			tasks[idx].Desc = desc
			tasks[idx].UpdatedAt = time.Now().Format(time.RFC3339)
		}
	}
	if err := writeList(tasks); err != nil {
		return fmt.Errorf("error with writing data.json: %v", err)
	}
	return nil
}

func TaskDeleteFunc(taskID string) error {
	oldTasks, err := readList()
	if err != nil {
		return fmt.Errorf("error with reading data.json: %v", err)
	}

	newTasks := []Task{}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("the id was expected to be a string: %v", err)
	}

	for _, task := range oldTasks {
		if task.ID != int64(id) {
			newTasks = append(newTasks, task)
		}
	}

	if err := writeList(newTasks); err != nil {
		return fmt.Errorf("error with writing data.json: %v", err)
	}
	return nil
}

func TaskMarkStatusFunc(taskID, status string) error {
	tasks, err := readList()
	if err != nil {
		return fmt.Errorf("error with reading data.json: %v", err)
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("the id was expected to be a string: %v", err)
	}

	for idx := range tasks {
		if tasks[idx].ID == int64(id) {
			tasks[idx].Status = status
			tasks[idx].UpdatedAt = time.Now().Format(time.RFC3339)
		}
	}
	if err := writeList(tasks); err != nil {
		return fmt.Errorf("error with writing data.json: %v", err)
	}
	return nil
}

func TaskListFunc(condition string) ([]Task, error) {
	tasks, err := readList()
	if err != nil {
		return nil, fmt.Errorf("error with reading data.json: %v", err)
	}

	if condition == "done" {
		resultTasks := []Task{}
		for _, task := range tasks {
			if task.Status == "done" {
				resultTasks = append(resultTasks, task)
			}
		}
		return resultTasks, nil
	} else if condition == "in-progress" {
		resultTasks := []Task{}
		for _, task := range tasks {
			if task.Status == "in progress" {
				resultTasks = append(resultTasks, task)
			}
		}
		return resultTasks, nil
	} else if condition == "todo" {
		resultTasks := []Task{}
		for _, task := range tasks {
			if task.Status == "todo" {
				resultTasks = append(resultTasks, task)
			}
		}
		return resultTasks, nil
	}
	return tasks, nil
}
