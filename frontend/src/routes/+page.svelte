<script lang="ts">
	import { Section } from 'flowbite-svelte-blocks';
	import {
		Accordion,
		AccordionItem,
		Button,
		CloseButton,
		Drawer,
		Hr,
		Input,
		Label,
		Select,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Textarea
	} from 'flowbite-svelte';
	import { EditSolid } from 'flowbite-svelte-icons';
	import { sineIn } from 'svelte/easing';

	import type { PageData } from './$types';
	import type { Template } from '$lib/backend/autoreplace';

	function replaceText(input: string, template: Template | undefined) {
		if (!template) return input;
		let output = input;
		for (const replacement of template.replacements) {
			output = output.replaceAll(replacement.from, replacement.to);
		}
		return output;
	}

	export let data: PageData;

	let selection = data.templates.map((_, id) => (id == 0 ? true : false));
	let content = '';

	$: selected = data.templates[0].id;
	$: selectedTemplate = data.templates.find((t) => t.id === selected);
	$: processedContent = replaceText(
		content,
		data.templates.find((t) => t.id === selected)
	);

	let hidden = true;
	let transitionParams = {
		x: -320,
		duration: 200,
		easing: sineIn
	};
	let categories = [
		{ value: '', name: 'Select category' },
		{ value: 'TV', name: 'TV/Monitors' },
		{ value: 'PC', name: 'PC' },
		{ value: 'GA', name: 'Gaming/Console' },
		{ value: 'PH', name: 'Phones' }
	];
	const handleCancel = () => {
		hidden = true;
	};
</script>

<div class="grid grid-flow-col grid-cols-3 grid-rows-3 gap-4">
	<div class="max-w-50 row-span-3">
		<Accordion>
			<!-- for each template, provide the corresponding application element -->
			{#each data.templates as template, index}
				<AccordionItem bind:open={selection[index]}>
					<span slot="header">
						{template.name}
						{#if template.id == selected}
							<b class="text-primary-500 dark:text-primary-400">active</b>
						{/if}
					</span>
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
					<div class="grid grid-flow-col justify-stretch">
						<Button class="mt-4 self-center" on:click={() => (selected = template.id)}>Use</Button>
						<div></div>
						<div></div>
						<div></div>
						<Button class="mt-4 place-content-end" on:click={() => (hidden = false)}>Edit</Button>
					</div>
				</AccordionItem>
			{/each}
		</Accordion>
	</div>
	<div class="col-span-2 row-span-2">
		<Label for="input" class="mb-2 font-semibold"
			>Text to process (using <b>{selectedTemplate.name}</b>):</Label
		>
		<Textarea id="input" placeholder="Text to process" rows="20" bind:value={content} />
		<Hr />
		<Label for="processed" class="mb-2 font-semibold">Text with replacements:</Label>
		<Textarea id="processed" placeholder="Processed text" rows="20" bind:value={processedContent} />
	</div>
</div>

<Section name="crudcreatedrawer">
	<Drawer transitionType="fly" {transitionParams} bind:hidden id="sidebar4">
		<div class="flex items-center">
			<h5
				id="drawer-label"
				class="mb-6 inline-flex items-center text-base font-semibold uppercase text-gray-500 dark:text-gray-400"
			>
				New Product
			</h5>
			<CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
		</div>
		<form action="#" class="mb-6">
			<div class="mb-6">
				<Label for="name" class="mb-2 block">Name</Label>
				<Input id="name" name="name" required placeholder="Type product name" />
			</div>
			<div class="mb-6">
				<Label for="bland" class="mb-2 block">Bland</Label>
				<Input id="bland" name="bland" required placeholder="Product brand" />
			</div>
			<div class="mb-6">
				<Label for="price" class="mb-2 block">Price</Label>
				<Input id="price" name="price" required placeholder="$2999" />
			</div>
			<div class="mb-6">
				<Label
					>Category
					<Select class="mt-2" items={categories} bind:value={selected} />
				</Label>
			</div>
			<div class="mb-6">
				<Label for="brand" class="mb-2">Description</Label>
				<Textarea id="message" placeholder="Enter event description here" rows="4" name="message" />
			</div>
			<div class="bottom-0 left-0 flex w-full justify-center space-x-4 pb-4 md:absolute md:px-4">
				<Button type="submit" class="w-full">Add product</Button>
				<Button class="w-full" color="light" on:click={handleCancel}>
					<svg
						aria-hidden="true"
						class="-ml-1 h-5 w-5 sm:mr-1"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/></svg
					>
					Cancel
				</Button>
			</div>
		</form>
	</Drawer>
</Section>
