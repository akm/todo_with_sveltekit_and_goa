import type { Todo } from '$lib/models/todo';
import { getTodos } from '$lib/server/grpc_client';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(event: ServerLoadEvent): Promise<{ todos: Todo[] }> {
	const { cookies } = event;
	let userid = cookies.get('userid');

	if (!userid) {
		userid = crypto.randomUUID();
		cookies.set('userid', userid, { path: '/' });
	}

	const todos = await getTodos();
	return { todos };
}
