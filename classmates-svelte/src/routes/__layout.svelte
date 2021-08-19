<script>
	import Nav from '../components/Nav.svelte';
</script>

<Nav />
<div class="flex flex-col items-center px-16">
	<slot />
</div>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
	
	:global(body) {
		background-color: #f2eee2;
		color: #0084f6;
		transition: background-color 0.3s
	}
	:global(body.dark-mode) {
		background-color: #1d3040;
		color: #bfc2c7;
	}
</style>