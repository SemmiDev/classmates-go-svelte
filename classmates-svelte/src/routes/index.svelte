<script>
	import StudentCard from '../components/StudentCard.svelte';
	import { students } from '../stores/studentstore';

	let searchInput = '';
	let filteredStudent = [];

	$: {
		if (searchInput) {
			filteredStudent = $students.filter((student) =>
				student.name.toLowerCase().includes(searchInput.toLowerCase())
			);
		} else {
			filteredStudent = [...$students];
		}
	}
</script>


<svelte:head>
	<title>Home</title>
</svelte:head>

<h1 class="text-4xl font-bold text-green-800 uppercase tracking-wide text-center pb-8">
	Classmates
</h1>

<input
	type="text"
	bind:value={searchInput}
	class="ring-1 border-green-400 rounded-md p-2 w-full focus:outline-none focus:ring-2 ring-green-400 focus:ring-offset-2"
	placeholder="Search your classmate"
/>
<div class="grid grid-cols-1 md:grid-cols-3 gap-3 py-4 w-full">
	{#each filteredStudent as student (student.id)}
		<StudentCard {student} />
	{/each}
</div>
