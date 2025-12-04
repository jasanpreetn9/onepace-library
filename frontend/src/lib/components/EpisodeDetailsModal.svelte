<!-- Episode Details Modal - Runes Mode
<script lang="ts">
	import { Dialog, DialogContent, DialogHeader, DialogTitle } from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Badge } from '$lib/components/ui/badge';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import { Download, ExternalLink } from 'lucide-svelte';

	let {
		open,
		onOpenChange,
		episode
	}: {
		open: boolean;
		onOpenChange: (open: boolean) => void;
		episode: UnifiedEpisode;
	} = $props();

	// export let open: boolean = false;
	// export let onOpenChange: (open: boolean) => void = () => {};

	// export let episodeNumber: number;
	// export let title: string;
	// export let description: string;
	// export let duration: string;
	// export let mangaChapters: string;
	// export let animeEpisodes: string;

	// export let status: "imported" | "missing" | "upgradable" | "downloading";

	// export type Version = {
	// 	id: string;
	// 	type: "normal" | "extended";
	// 	resolution: string;
	// 	size: string;
	// 	seeders: number;
	// };

	// export let installedVersion: Version | null = null;
	// export let availableVersions: Version[] = [];

	const statusColors: Record<'none' | 'imported' | 'missing' | 'upgradable' | 'queued' | 'downloading' | 'failed', string> = {
		none: 'bg-muted text-muted-foreground',
		imported: 'bg-success text-success-foreground',
		missing: 'bg-destructive text-destructive-foreground',
		upgradable: 'bg-warning text-warning-foreground',
		queued: 'bg-primary/50 text-primary-foreground',
		downloading: 'bg-primary text-primary-foreground',
		failed: 'bg-destructive/70 text-destructive-foreground'
	};
</script>

<Dialog {open} {onOpenChange}>
	<DialogContent class="max-w-3xl">
		<DialogHeader>
			<div class="flex items-center gap-3">
				<DialogTitle class="text-xl">
					Episode {episode.episode}: {episode.title}
				</DialogTitle>

				<Badge class={statusColors[episode.downloadStatus]}>
					{episode.downloadStatus.charAt(0).toUpperCase() + episode.downloadStatus.slice(1)}
				</Badge>
			</div>
		</DialogHeader>

		<ScrollArea class="max-h-[70vh] pr-4">
			<div class="space-y-6">
				<div class="grid grid-cols-3 gap-4 rounded-lg bg-muted/50 p-4">
					<div>
						<p class="text-xs font-medium text-muted-foreground">Released</p>
						<p class="mt-1 text-sm font-semibold">{episode.released}</p>
					</div>

					 <div>
						<p class="text-xs font-medium text-muted-foreground">Manga Chapters</p>
						<p class="mt-1 text-sm font-semibold">{episode.}</p>
					</div>

					<div>
						<p class="text-xs font-medium text-muted-foreground">Anime Episodes</p>
						<p class="mt-1 text-sm font-semibold">{animeEpisodes}</p>
					</div>
				</div>

				<div>
					<h3 class="mb-2 text-sm font-semibold">Description</h3>
					<p class="text-sm leading-relaxed text-muted-foreground">
						{episode.description}
					</p>
				</div>

				{#if episode.version}
					<div>
						<h3 class="mb-3 text-sm font-semibold">Installed Version</h3>

						<div class="rounded-lg border bg-card p-4 flex items-center justify-between">
							<div class="flex items-center gap-3">
								<Badge variant="outline" class="capitalize">
									{installedVersion.type}
								</Badge>

								<span class="text-sm">{installedVersion.resolution}</span>
								<span class="text-sm text-muted-foreground">{installedVersion.size}</span>
							</div>

							<Button size="sm" variant="outline">
								<ExternalLink class="mr-2 h-4 w-4" />
								View File
							</Button>
						</div>
					</div>
				{/if} -->

				<!-- AVAILABLE VERSIONS -->
				<!-- <div>
					<h3 class="mb-3 text-sm font-semibold">Available Versions</h3>

					<div class="space-y-2">
						{#each availableVersions as v}
							<div
								class="flex items-center justify-between rounded-lg border bg-card p-4 transition hover:border-primary/50"
							>
								<div class="flex items-center gap-4">
									<Badge variant="outline" class="capitalize">{v.type}</Badge>

									<span class="text-sm font-medium">{v.resolution}</span>
									<span class="text-sm text-muted-foreground">{v.size}</span>
									<span class="text-sm text-muted-foreground">{v.seeders} seeders</span>
								</div>

								<Button size="sm" on:click={() => dispatch('download', v)}>
									<Download class="mr-2 h-4 w-4" />
									Download
								</Button>
							</div>
						{/each}
					</div>
				</div>
			</div>
		</ScrollArea>
	</DialogContent>
</Dialog>

<style>
	/* Optional: Improve smooth scrolling inside modal */
	:global(.scroll-area-viewport) {
		scrollbar-width: thin;
	}
</style> -->

<script lang="ts">
	import {
		Dialog,
		DialogContent,
		DialogHeader,
		DialogTitle
	} from "$lib/components/ui/dialog";

	import { Badge } from "$lib/components/ui/badge";
	import { ScrollArea } from "$lib/components/ui/scroll-area";

	export let open: boolean;
	export let onOpenChange: (open: boolean) => void;
	export let episode: UnifiedEpisode | null;

	const statusColors = {
		imported: "bg-success text-success-foreground",
		missing: "bg-destructive text-destructive-foreground",
		upgradable: "bg-warning text-warning-foreground",
		downloading: "bg-primary text-primary-foreground",
		queued: "bg-primary/60 text-white",
		failed: "bg-red-700 text-white",
		none: "bg-muted text-muted-foreground"
	};
</script>

<Dialog {open} {onOpenChange}>
	<DialogContent class="max-w-3xl">
		{#if episode}
			<DialogHeader>
				<div class="flex items-center gap-3">
					<DialogTitle class="text-xl">
						Episode {episode.episode}: {episode.title}
					</DialogTitle>

					<Badge class={statusColors[episode.downloadStatus]}>
						{episode.downloadStatus}
					</Badge>
				</div>
			</DialogHeader>

			<ScrollArea class="max-h-[70vh] pr-4">
				<div class="space-y-6">
					<p class="text-muted-foreground">{episode.description}</p>

					<div class="grid grid-cols-3 gap-4 bg-muted/50 p-4 rounded-lg">
						<div>
							<p class="text-xs text-muted-foreground">Released</p>
							<p class="font-medium">{episode.released}</p>
						</div>

						{#if episode.filePath}
							<div>
								<p class="text-xs text-muted-foreground">File</p>
								<p class="font-mono text-xs">{episode.filePath}</p>
							</div>
						{/if}
					</div>
				</div>
			</ScrollArea>
		{/if}
	</DialogContent>
</Dialog>
