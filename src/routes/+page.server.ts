import { getTodos } from '$lib/server/client';
import type { Todo } from '$lib/models/todo';
import type { ServerLoadEvent } from '@sveltejs/kit';

export const ssr = false;

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
