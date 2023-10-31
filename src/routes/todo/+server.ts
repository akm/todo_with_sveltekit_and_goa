import { json, type RequestEvent } from '@sveltejs/kit';
// import * as database from '$lib/server/database';
import { createTodo } from '$lib/server/grpc_client';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function POST(event: RequestEvent) {
	const { request, cookies } = event;
	const { description } = await request.json();
	if (!(typeof description === 'string')) throw new Error('Description not found');

	const userid = cookies.get('userid');
	if (!userid) throw new Error('User not found');
	const { id } = await createTodo({ description });

	return json({ id }, { status: 201 });
}
