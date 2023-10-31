import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

import type { Todo } from '$lib/models/todo';
import { TodosClient, type ITodosClient } from './goagen_goa_todo_example_todos.client';

const transport = new GrpcTransport({
	host: 'localhost:8080',
	channelCredentials: ChannelCredentials.createInsecure()
});

const client: ITodosClient = new TodosClient(transport);

export const getTodos = async (): Promise<Todo[]> => {
	const { response } = await client.list({});
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id.toString(),
				description: item.title,
				done: item.state === 'closed'
			};
		}) ?? []
	);
};

export const createTodo = async (arg: { description: string }): Promise<{ id: string }> => {
	const { description } = arg;
	const { response } = await client.create({
		title: description,
		state: 'open'
	});
	return { id: response.id.toString() };
};

export const toggleTodo = async (arg: { id: string | undefined; done: boolean }): Promise<void> => {
	if (!arg.id) return;
	const { response } = await client.show({ id: BigInt(arg.id) });
	await client.update({
		id: BigInt(arg.id),
		title: response.title,
		state: arg.done ? 'closed' : 'open'
	});
};

export const deleteTodo = async (arg: { id: string | undefined }): Promise<void> => {
	if (!arg.id) return;
	await client.delete({ id: BigInt(arg.id) });
};
