<script lang="ts">
	import '../app.css';
	import ToastContainer from '$lib/toast/ToastContainer.svelte';
</script>

<main class="min-h-screen bg-base-200">
	<slot />
</main>

<!-- Global Toast Container -->
<ToastContainer />