<script>
	export let std;
</script>

<div class="bg-green-400 p-2 rounded-md flex flex-col items-center">
	<img src={std.image_url} alt={std.name} class="h-48 bg-white rounded-md" />
	<div class="rounded-md p-2 w-full bg-white mt-2 space-y-2">
		<div class="flex justify-between w-full">
			<span class="font-semibold">Name</span>
			<span>{std.name}</span>
		</div>
		<div class="flex justify-between w-full">
			<span class="font-semibold">Identifier</span>
			<span>{std.identifier}</span>
		</div>
		<div class="flex justify-between w-full">
			<span class="font-semibold">Email</span>
			<span>{std.email}</span>
		</div>
	</div>
</div>
