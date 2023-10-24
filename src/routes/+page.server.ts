import { getTodos } from '$lib/apisvr/client';
import type { Todo } from '$lib/apisvr/client';
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
