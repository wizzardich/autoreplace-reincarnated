<script lang="ts">
	import '../app.pcss';
	import { page } from '$app/stores';
	import { DarkMode, Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';

	let width: number;
	$: activeUrl = $page.url.pathname;

	let divClass = 'w-full ml-auto lg:block lg:w-auto order-1 lg:order-none';
	let ulClass =
		'flex flex-col py-3 my-4 lg:flex-row lg:my-0 text-sm font-medium gap-4 dark:lg:bg-transparent lg:bg-white lg:border-0';
</script>

<svelte:window bind:innerWidth={width} />
<header class="mx-auto w-full flex-none bg-white dark:bg-slate-950">
	<Navbar let:hidden let:toggle>
		<NavBrand href="/" class="lg:ml-32">
			<span class="self-center whitespace-nowrap pl-4 text-xl font-semibold dark:text-white">
				AutoReplace <sup class="text-primary-500 dark:text-primary-400 text-xs">reincarnated</sup>
			</span>
		</NavBrand>
		<NavUl
			{hidden}
			{divClass}
			{ulClass}
			nonActiveClass="md:!pl-3 md:!py-2 lg:!pl-0 text-gray-700 hover:bg-gray-100 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 dark:text-white lg:dark:hover:text-primary-700 dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent"
			activeClass="md:!pl-3 md:!py-2 lg:!pl-0 lg:text-primary-700 text-white dark:text-white dark:lg:text-primary-500 bg-primary-700 lg:bg-transparent dark:bg-primary-600 lg:dark:bg-transparent cursor-default"
		>
			<NavLi class="lg:mb-0 lg:px-2" active={activeUrl === '/'} href="/">Home</NavLi>
			<NavLi class="lg:mb-0 lg:px-2" href="https://plots.smugglersden.org">Plots</NavLi>
			<NavLi class="lg:mb-0 lg:px-2" href="https://github.com/wizzardich/autoreplace-reincarnated"
				>GitHub
			</NavLi>
		</NavUl>
		<div class="ml-auto flex lg:mr-16">
			<DarkMode class="inline-block hover:text-gray-900 dark:hover:text-white" />
		</div>
	</Navbar>
</header>

<div class="mx-auto flex w-full px-40 pr-40">
	<main class="mx-auto w-full">
		<slot />
	</main>
</div>
