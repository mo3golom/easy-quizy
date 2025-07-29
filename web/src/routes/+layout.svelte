<script lang="ts">
	import "../app.css";
	import ToastContainer from "$lib/toast/ToastContainer.svelte";
	import { init, isTMA } from "@telegram-apps/sdk";
	import { onMount } from "svelte";
	import { initData } from "@telegram-apps/sdk";
	import { setTelegramUser, restoreUser } from "$lib/stores/user";
	import Loading from "$lib/components/Loading.svelte";
	interface Props {
		children?: import('svelte').Snippet;
	}

	let { children }: Props = $props();

	let isLoading: boolean = $state(true);

	onMount(async () => {
		restoreUser();

		try {
			if (isTMA()) {
				await init();

				initData.restore();

				const telegramUser = initData.user();
				if (telegramUser) {
					setTelegramUser(telegramUser);
				}
			}

			isLoading = false;
		} catch (error) {
			console.error("Error initializing Telegram SDK:", error);
		}
	});
</script>

<main class="min-h-screen bg-primary-content relative">
	{#if isLoading}
		<Loading />
	{:else}
		{@render children?.()}
	{/if}
</main>

<!-- Global Toast Container -->
<ToastContainer />
