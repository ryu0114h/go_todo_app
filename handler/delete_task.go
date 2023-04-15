package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ryu0114h/go_todo_app/entity"
)

type DeleteTask struct {
	Service DeleteTaskService
}

func (dt *DeleteTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.Split(r.URL.Path, "/tasks/")[1]
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	err = dt.Service.DeleteTask(r.Context(), entity.TaskID(idInt))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
