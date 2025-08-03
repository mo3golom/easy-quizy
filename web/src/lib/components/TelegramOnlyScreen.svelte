<script lang="ts">
    import { env } from "$env/dynamic/public";
</script>

<div class="w-full min-h-screen bg-primary flex items-center">
    <div class="container mx-auto max-w-md p-2">
        <div class="card p-4 bg-primary-content rounded-3xl">
            <div class="card-body text-center">
                <div class="text-6xl mb-4">📱</div>
                <h1 class="card-title text-2xl mb-4 justify-center">
                    Telegram Mini App
                </h1>
                <p class="text-base-content/70 mb-6">
                    Это приложение можно открыть только внутри Telegram Mini App
                </p>
                <div class="card-actions justify-center">
                    <a
                        href={env.PUBLIC_TG_BOT_URL}
                        class="btn btn-secondary transition-transform"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        Открыть в Telegram
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>
