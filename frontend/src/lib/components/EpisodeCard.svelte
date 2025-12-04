<!-- Episode Card - Runes Mode -->
<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Badge } from '$lib/components/ui/badge';
	import { Eye, Download, MoreHorizontal, CheckCircle2 } from 'lucide-svelte';
	import { cn } from '$lib/utils';
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	function select() {
		dispatch("select", episode);
	}
	let {
		episode,
		onClick = () => {},
		onToggleMonitor = () => {},
		onDownload = () => {}
	}: {
		episode: UnifiedEpisode;
		onClick?: () => void;
		onToggleMonitor?: () => void;
		onDownload?: () => void;
	} = $props();
</script>

<div
	class="group border border-border bg-card rounded-lg p-5 hover:border-primary/50 hover:shadow-md transition-all"
	onclick={select}
>
	<div class="flex items-start justify-between gap-4">
		<div class="flex-1 space-y-2">
			<div class="flex items-center gap-3">
				<span class="text-sm font-medium text-muted-foreground">
					Episode {episode.episode}
				</span>

				<h3 class="text-base font-semibold">{episode.title}</h3>

				<!-- Status Badge -->
				<Badge
					class={cn(
						status === 'imported' && 'bg-green-600 text-white',
						status === 'missing' && 'bg-red-600 text-white',
						status === 'upgradable' && 'bg-yellow-600 text-white'
					)}
				>
					{status}
				</Badge>
			</div>

			<p class="text-sm text-muted-foreground">
				{episode.description || 'No description available.'}
			</p>

			{#if episode.version || episode.crc32}
				<div class="flex items-center gap-4 text-xs text-muted-foreground">
					{#if episode.version}<span>Version: {episode.version}</span>{/if}
					{#if episode.crc32}<span>CRC: {episode.crc32}</span>{/if}
				</div>
			{/if}

			{#if episode.filePath}
				<div class="rounded bg-muted/50 px-3 py-2">
					<p class="font-mono text-xs text-muted-foreground">{episode.filePath}</p>
				</div>
			{/if}
		</div>

		<!-- Actions -->
		<div class="flex flex-col items-end gap-2">
			<Button size="sm" variant={episode.monitored ? 'default' : 'outline'} onclick={onToggleMonitor}>
				<CheckCircle2 class="mr-2 h-4 w-4" />
				{episode.monitored ? 'Monitored' : 'Unmonitored'}
			</Button>

			{#if status === 'missing'}
				<Button size="sm" variant="outline" onclick={onDownload}>
					<Download class="mr-2 h-4 w-4" />
					Download
				</Button>
			{/if}

			{#if status === 'upgradable'}
				<Button size="sm" variant="outline" onclick={onDownload}>
					<Download class="mr-2 h-4 w-4" />
					Upgrade
				</Button>
			{/if}

			<Button size="sm" variant="ghost" onclick={onClick}>
				<Eye class="mr-2 h-4 w-4" />
				Details
			</Button>

			<Button size="icon" variant="ghost">
				<MoreHorizontal class="h-4 w-4" />
			</Button>
		</div>
	</div>
</div>
