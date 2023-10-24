<script lang="ts">
	import { createTodo, toggleTodo, deleteTodo } from '$lib/apisvr/client';
	import type { Todo } from '$lib/apisvr/client';
	export let data: { todos: Todo[]};
</script>

<div class="centered">
	<h1>todos</h1>

	<label>
		add a todo:
		<input
			type="text"
			autocomplete="off"
			on:keydown={async (e) => {
				if (e.key === 'Enter') {
					const input = e.currentTarget;
					const description = input.value;

					const todo = await createTodo({description});
					const { id } = todo;

					data.todos = [...data.todos, {id, description, done: false}];

					input.value = '';
				}
			}}
		/>
	</label>

	<ul class="todos">
		{#each data.todos as todo (todo.id)}
			<li>
				<label>
					<input
						type="checkbox"
						checked={todo.done}
						on:change={async (e) => {
							await toggleTodo({id: todo.id, done: e.currentTarget.checked})
						}}
					/>
					<span>{todo.description}</span>
					<button
						aria-label="Mark as complete"
						on:click={async (e) => {
							await deleteTodo({id: todo.id})
							data.todos = data.todos.filter((t) => t !== todo);
						}}
					/>
				</label>
			</li>
		{/each}
	</ul>
</div>

<style>
	.centered {
		max-width: 20em;
		margin: 0 auto;
	}

	label {
		display: flex;
		width: 100%;
	}

	input[type="text"] {
		flex: 1;
		color: black;
	}

	span {
		flex: 1;
	}

	button {
		border: none;
		background: url(./remove.svg) no-repeat 50% 50%;
		background-size: 1rem 1rem;
		cursor: pointer;
		height: 100%;
		aspect-ratio: 1;
		opacity: 0.5;
		transition: opacity 0.2s;
	}

	button:hover {
		opacity: 1;
	}
</style>
	