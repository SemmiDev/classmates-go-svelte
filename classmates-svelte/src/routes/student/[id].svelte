<script context="module">
	export async function load({ page }) {
		const id = page.params.id; 
		const u = 'http://localhost:8080/api/v1/classmates/${id}';
		const res = await fetch(u);
		const std = await res.json();

		return {
			props: {
				std
			}
		};
	}
</script>

<script>
	import StudentDetails from '../../components/StudentDetailsCard.svelte';
	export let std;
</script>


<StudentDetails {std} />
