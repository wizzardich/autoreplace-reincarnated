<script lang="ts">
	import {
		AccordionItem,
		Accordion,
		Label,
		Textarea,
		Button,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Hr
	} from 'flowbite-svelte';
	import type { PageData } from './$types';
	import type { Template } from '$lib/backend/autoreplace';

	export let data: PageData;
	let selection = data.templates.map((_, id) => (id == 0 ? true : false));
	$: selected = data.templates[0].id;
	let content = '';
	$: processedContent = replaceText(
		content,
		data.templates.find((t) => t.id === selected)
	);

	function replaceText(input: string, template: Template | undefined) {
		console.log('input', input);
		if (!template) return input;
		let output = input;
		for (const replacement of template.replacements) {
			output = output.replaceAll(replacement.from, replacement.to);
		}
		return output;
	}
</script>

<div class="grid grid-flow-col grid-cols-3 grid-rows-3 gap-4">
	<div class="max-w-50 row-span-3">
		<Accordion>
			<!-- for each template, provide the corresponding application element -->
			{#each data.templates as template, index}
				<AccordionItem bind:open={selection[index]}>
					<span slot="header">{template.name}</span>
					<Table striped={true}>
						<TableHead>
							<TableHeadCell>Match</TableHeadCell>
							<TableHeadCell>Replace with</TableHeadCell>
						</TableHead>
						<TableBody tableBodyClass="divide-y">
							{#each template.replacements as replacement}
								<TableBodyRow>
									<TableBodyCell>
										{replacement.from}
									</TableBodyCell>
									<TableBodyCell>
										{replacement.to}
									</TableBodyCell>
								</TableBodyRow>
							{/each}
						</TableBody>
					</Table>
					<Button class="mt-4 self-center" on:click={() => (selected = template.id)}>Replace</Button
					>
				</AccordionItem>
			{/each}
		</Accordion>
	</div>
	<div class="col-span-2 row-span-2">
		<Label for="input-text" class="mb-2 font-semibold">Text to process</Label>
		<Textarea id="input-text" placeholder="Text to process" rows="20" bind:value={content} />
		<Hr />
		<Label for="processed-text" class="mb-2 font-semibold">Text with replacements:</Label>
		<Textarea
			id="processed-text"
			placeholder="Processed text"
			rows="20"
			bind:value={processedContent}
		/>
	</div>
</div>
