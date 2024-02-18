<script lang="ts">
	import { Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import type { PageData } from './$types';
	import { Alert } from 'flowbite-svelte';

	// @ts-ignore
	function handleClick(t: string) {
		return (e: MouseEvent) => {
			e.preventDefault();
			alert('Clicked on: ' + t);
		};
	}

	export let data: PageData;
</script>

<div class="p-8">
	<Alert>
		<span class="font-medium">Info alert!</span>
		Change a few things up and try submitting again.
	</Alert>
</div>

<Button>Dropdown button</Button>
<Dropdown>
	{#each data.templates as template}
		<DropdownItem on:click={handleClick(template.id)}>{template.name}</DropdownItem>
	{/each}
</Dropdown>
