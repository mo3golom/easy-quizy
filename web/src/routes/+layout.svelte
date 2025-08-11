<script lang="ts">
	import "../app.css";
	import ToastContainer from "$lib/toast/ToastContainer.svelte";
	import { init, isTMA } from "@telegram-apps/sdk";
	import { onMount } from "svelte";
	import { initData } from "@telegram-apps/sdk";
	import { setTelegramUser, restoreUser} from "$lib/stores/user";
	import Loading from "$lib/components/Loading.svelte";
	import TelegramOnlyScreen from "$lib/components/TelegramOnlyScreen.svelte";
	import { goto } from "$app/navigation";
	import { mockEnvironment} from "$lib/utils/telegram"

	interface Props {
		children?: import("svelte").Snippet;
	}

	let { children }: Props = $props();
	let isLoading: boolean = $state(true);
	let isTelegramApp: boolean = $state(false);

	onMount(async () => {
		mockEnvironment();
		restoreUser();

		try {
			isTelegramApp = isTMA();

			if (isTelegramApp) {
				await init();

				initData.restore();

				const telegramUser = initData.user();
				const telegramUserChatID = initData.chatInstance();
				const telegramUserChatType = initData.chatType();
				if (telegramUser) {
					if (telegramUserChatID && telegramUserChatType) {
						setTelegramUser(
							telegramUser,
							Number(telegramUserChatID),
							telegramUserChatType,
						);
					} else {
						setTelegramUser(telegramUser);
					}
				}

				const startParam = initData.startParam();
				if (startParam) {
					try {
						const decodedString = atob(startParam);
						const parsedData = JSON.parse(decodedString);

						if (
							parsedData.toPage &&
							parsedData.toPage.trim() !== ""
						) {
							await goto(parsedData.toPage);
						}
					} catch (error) {
						// Ignore any parsing or decoding errors
					}
				}
			}
		} catch (error) {
			console.error("Error initializing Telegram SDK:", error);
		} finally {
			isLoading = false;
		}
	});
</script>

<main class="min-h-screen bg-primary-content relative">
	{#if isLoading}
		<Loading />
	{:else if !isTelegramApp}
		<TelegramOnlyScreen />
	{:else}
		{@render children?.()}
	{/if}
</main>

<!-- Global Toast Container -->
<ToastContainer />
