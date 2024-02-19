<script lang="ts">
	import { Select, Label, Textarea } from 'flowbite-svelte';
	import type { ActionData, PageData } from './$types';

	export let data: PageData;
	let options = data.templates.map((t) => ({ value: t.id, name: t.name }));

	function getTemplateContent(id: string) {
		let template = data.templates.find((t) => t.id === id);
		console.log(data.templates);
		if (!template) return '';

		let reprMap = template.replacements.map((r) => r.from + ' -- ' + r.to);
		return reprMap.join('\n');
	}

	let selected = options[0].value;
	$: content = getTemplateContent(selected);
</script>

<div class="centered">
	<form action="?/loadTemplate" method="POST">
		<Label for="template-selector">Select an option</Label>
		<Select name="template-id" class="mt-2" items={options} bind:value={selected} />
		<br />
		<Label for="current-template">Current Template</Label>
		<Textarea id="current-template" placeholder="Current Template" rows="4" bind:value={content} />
	</form>
</div>

<style>
	.centered {
		max-width: 20em;
		margin: 0 auto;
	}
</style>
