<script lang="ts">
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { backButton, isTMA } from "@telegram-apps/sdk";

    interface Props {
        targetPage: string;
    }

    let { targetPage }: Props = $props();
    let showFallback: boolean = $state(false);

    async function handleClick() {
        if (targetPage) {
            await goto(targetPage);
        }
    }

    onMount(() => {
        if (isTMA() && backButton.mount.isAvailable()) {
            backButton.mount();
            backButton.onClick(handleClick);
            backButton.show();
        } else {
            showFallback = true;
        }

        return () => {
            if (backButton.isMounted()) {
                backButton.hide();
                backButton.offClick(handleClick);
                backButton.unmount();
            }
        };
    });
</script>

{#if showFallback}
    <div>
        <button class="btn btn-ghost btn-sm flex items-center gap-2" onclick={handleClick}>
            <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="size-4"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M10.5 19.5 3 12m0 0 7.5-7.5M3 12h18"
                />
            </svg>
            назад
        </button>
    </div>
{/if}
