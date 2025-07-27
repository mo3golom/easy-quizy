<script lang="ts">
	import '../app.css';
	import ToastContainer from '$lib/toast/ToastContainer.svelte';
</script>

<main class="min-h-screen bg-primary-content">
	<slot />
</main>

<!-- Global Toast Container -->
<ToastContainer />