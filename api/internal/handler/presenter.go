package handler

import (
	"fmt"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
)

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

type TaskResponse struct {
	*entity.Task
	Links []Link `json:"links"`
}

func BuildTaskResponse(task *entity.Task, baseURL string) TaskResponse {
	taskURL := fmt.Sprintf("%s/tasks/%s", baseURL, task.ID)

	links := []Link{
		{Rel: "self", Href: taskURL, Method: "GET"},
	}

	// Regra dinâmica do HATEOAS:
	// Só enviamos os links de alteração se a tarefa NÃO estiver concluída.
	if task.Status != entity.StatusDone {
		links = append(links, Link{
			Rel: "update", Href: taskURL, Method: "PUT",
		})
		links = append(links, Link{
			Rel: "update_status", Href: taskURL + "/status", Method: "PATCH",
		})
		links = append(links, Link{
			Rel: "delete", Href: taskURL, Method: "DELETE",
		})
	}

	return TaskResponse{
		Task:  task,
		Links: links,
	}
}
