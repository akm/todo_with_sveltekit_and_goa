import { promisify } from 'util';
import { credentials } from '@grpc/grpc-js';

import type { Todo } from '$lib/models/todo';
import { TodosClient, type ITodosClient } from './goagen_goa_todo_example_todos_grpc_pb';
import {
	ListRequest,
	ListResponse,
	CreateRequest,
	CreateResponse,
	UpdateRequest,
	UpdateResponse,
	DeleteRequest,
	DeleteResponse
} from './goagen_goa_todo_example_todos_pb';

const client: ITodosClient = new TodosClient(`localhost:8080`, credentials.createInsecure());

export const getTodos = async (): Promise<Todo[]> => {
	const request = new ListRequest();
	const response = await promisify<ListRequest, ListResponse>(client.list).bind(client)(request);

	return (
		response
			.getItems()
			?.getFieldList()
			.map((item) => {
				return {
					id: item.getId().toString(),
					description: item.getTitle(),
					done: item.getState() === 'closed'
				};
			}) ?? []
	);
};

export const createTodo = async (arg: { description: string }): Promise<{ id: string }> => {
	const { description } = arg;

	const request = new CreateRequest();
	request.setTitle(description);

	const response = await promisify<CreateRequest, CreateResponse>(client.create).bind(client)(
		request
	);
	return { id: response.getId().toString() };
};

export const toggleTodo = async (arg: { id: string | undefined; done: boolean }): Promise<void> => {
	if (!arg.id) return;
	const request = new UpdateRequest();
	request.setId(parseInt(arg.id));
	request.setState(arg.done ? 'closed' : 'open');
	await promisify<UpdateRequest, UpdateResponse>(client.update).bind(client)(request);
};

export const deleteTodo = async (arg: { id: string | undefined }): Promise<void> => {
	if (!arg.id) return;
	const request = new DeleteRequest();
	request.setId(parseInt(arg.id));
	await promisify<DeleteRequest, DeleteResponse>(client.delete).bind(client)(request);
};
