<script>
	export let student;
</script>

<a
	href={`/student/${student.id}`}
	class="rounded-md bg-green-50 flex flex-col items-center justify-center cursor-pointer transition hover:shadow-lg py-5">
	<img src={student.image} alt={student.name} class="h-48" />
	<p class="uppercase tracking-wide font-medium text-green-900">{student.name}</p>
</a>
