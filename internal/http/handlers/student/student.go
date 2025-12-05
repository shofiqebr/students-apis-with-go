package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shofiqebr/students-apis/internal/storage"
	"github.com/shofiqebr/students-apis/internal/types"
	"github.com/shofiqebr/students-apis/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		slog.Info("creating a student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return 
		}

		if err != nil {
			response.WriteJson(w,http.StatusBadRequest, response.GeneralError(err))
			return                                                                                                                         
		}

		// request validation

		if err := validator.New().Struct(student); err != nil {
			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest,response.ValidationError(validateError))
			return 
		}

		
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}