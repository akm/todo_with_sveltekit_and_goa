// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	todos "github.com/akm/goa_todo_example/gen/todos"
	todosviews "github.com/akm/goa_todo_example/gen/todos/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "todos" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListTodosPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todos", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the todos
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todos", "list", err)
			}
			p := NewListTodoListOK(&body)
			view := "default"
			vres := &todosviews.TodoList{Projected: p, View: view}
			if err = todosviews.ValidateTodoList(vres); err != nil {
				return nil, goahttp.ErrValidationError("todos", "list", err)
			}
			res := todos.NewTodoList(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todos", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "todos" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*todos.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todos", "show", "*todos.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowTodosPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todos", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the todos
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todos", "show", err)
			}
			p := NewShowTodoOK(&body)
			view := "default"
			vres := &todosviews.Todo{Projected: p, View: view}
			if err = todosviews.ValidateTodo(vres); err != nil {
				return nil, goahttp.ErrValidationError("todos", "show", err)
			}
			res := todos.NewTodo(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todos", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "todos" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateTodosPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todos", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the todos create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todos.TodoCreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("todos", "create", "*todos.TodoCreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todos", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the todos
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todos", "create", err)
			}
			p := NewCreateTodoCreated(&body)
			view := "default"
			vres := &todosviews.Todo{Projected: p, View: view}
			if err = todosviews.ValidateTodo(vres); err != nil {
				return nil, goahttp.ErrValidationError("todos", "create", err)
			}
			res := todos.NewTodo(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todos", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "todos" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*todos.TodoUpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todos", "update", "*todos.TodoUpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateTodosPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todos", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the todos update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todos.TodoUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("todos", "update", "*todos.TodoUpdatePayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todos", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the todos
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todos", "update", err)
			}
			p := NewUpdateTodoOK(&body)
			view := "default"
			vres := &todosviews.Todo{Projected: p, View: view}
			if err = todosviews.ValidateTodo(vres); err != nil {
				return nil, goahttp.ErrValidationError("todos", "update", err)
			}
			res := todos.NewTodo(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todos", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "todos" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*todos.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todos", "delete", "*todos.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteTodosPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todos", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteResponse returns a decoder for responses returned by the todos
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todos", "delete", err)
			}
			p := NewDeleteTodoOK(&body)
			view := "default"
			vres := &todosviews.Todo{Projected: p, View: view}
			if err = todosviews.ValidateTodo(vres); err != nil {
				return nil, goahttp.ErrValidationError("todos", "delete", err)
			}
			res := todos.NewTodo(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todos", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTodoListItemResponseBodyToTodosviewsTodoListItemView builds a value
// of type *todosviews.TodoListItemView from a value of type
// *TodoListItemResponseBody.
func unmarshalTodoListItemResponseBodyToTodosviewsTodoListItemView(v *TodoListItemResponseBody) *todosviews.TodoListItemView {
	res := &todosviews.TodoListItemView{
		ID:        v.ID,
		Title:     v.Title,
		State:     v.State,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}

	return res
}
