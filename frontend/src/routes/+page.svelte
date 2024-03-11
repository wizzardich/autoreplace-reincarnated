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
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Textarea
	} from 'flowbite-svelte';
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

	let selection = data.templates.map((t, id) => (id == 0 ? true : false));
	let content = '';

	$: selected = data.templates[0].id;
	$: selectedTemplate = data.templates.find((t) => t.id === selected);
	$: processedContent = replaceText(
		content,
		data.templates.find((t) => t.id === selected)
	);

	let isEditMenuHidden = true;
	let editableTemplate = data.templates[0];
	let transitionParams = {
		x: -320,
		duration: 200,
		easing: sineIn
	};

	const handleEdit = (t: Template) => {
		isEditMenuHidden = false;
		editableTemplate = t;
	};
	const handleCancel = () => {
		isEditMenuHidden = true;
	};

	async function onTemplateEdit(event: Event) {
		event.preventDefault();
		console.log('Template edited', editableTemplate);
		const response = await fetch(`/templates/edit`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(editableTemplate)
		});
		if (response.ok) {
			console.log('Template edited successfully');
			data.templates = data.templates.map((t) =>
				t.id === editableTemplate.id ? editableTemplate : t
			);
			isEditMenuHidden = true;
		} else {
			console.error('Failed to edit template');
		}
	}
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
						<Button class="mt-4 place-content-end" on:click={() => handleEdit(template)}
							>Edit</Button
						>
					</div>
				</AccordionItem>
			{/each}
		</Accordion>
	</div>
	<div class="col-span-2 row-span-2">
		<Label for="input" class="mb-2 font-semibold"
			>Text to process{#if selectedTemplate}
				(using <b>{selectedTemplate.name}</b>){/if}:</Label
		>
		<Textarea id="input" placeholder="Text to process" rows="20" bind:value={content} />
		<Hr />
		<Label for="processed" class="mb-2 font-semibold">Text with replacements:</Label>
		<Textarea id="processed" placeholder="Processed text" rows="20" bind:value={processedContent} />
	</div>
</div>

<Section name="crudcreatedrawer">
	<Drawer
		transitionType="fly"
		{transitionParams}
		bind:hidden={isEditMenuHidden}
		id="edit-sidebar"
		width="w-160"
	>
		<div class="flex items-center">
			<h5
				id="drawer-label"
				class="mb-6 inline-flex items-center text-base font-semibold uppercase text-gray-500 dark:text-gray-400"
			>
				Edit Replacement Template
			</h5>
			<CloseButton on:click={() => (isEditMenuHidden = true)} class="mb-4 dark:text-white" />
		</div>
		<form action="edit" on:submit|preventDefault={onTemplateEdit} class="mb-6">
			<div class="mb-6">
				<Label for="name" class="mb-2 block">Name</Label>
				<Input
					id="name"
					name="name"
					required
					placeholder="Type product name"
					bind:value={editableTemplate.name}
				/>
			</div>
			{#each editableTemplate.replacements as replacement, _}
				<div>
					<!--<div class="flex items-center justify-between">-->
					<div class="flex w-full items-center justify-between space-x-4 pb-4">
						<Label for="from" class="block">From</Label>
						<Label for="to" class="block">To</Label>
					</div>
					<div class="flex w-full justify-center space-x-4 pb-4">
						<Input
							id="from"
							name="from"
							required
							placeholder="Type text to replace"
							bind:value={replacement.from}
						/>
						<Input
							id="to"
							name="to"
							required
							placeholder="Type replacement text"
							bind:value={replacement.to}
						/>
					</div>
				</div>
			{/each}
			<div class="bottom-0 left-0 flex w-full justify-center space-x-4 pb-4">
				<Button type="submit" class="w-full">Save Changes</Button>
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
