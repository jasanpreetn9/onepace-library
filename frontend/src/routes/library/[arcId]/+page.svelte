<!-- [arcId]/+page.svelte -->
<script lang="ts">
	import EpisodeCard from '$lib/components/EpisodeCard.svelte';
	import EpisodeDetailsModal from '$lib/components/EpisodeDetailsModal.svelte';

	import { page } from '$app/state';
	import { arcs } from '$lib/stores';

	const arcId = $derived(page.params.arcId);

	let arcData = $state<UnifiedArc | null>(null);
	let selectedEpisode = $state<UnifiedEpisode | null>(null);

	$effect(() => {
		arcData = $arcs.find(a => a.arc.toString() === arcId) ?? null;
	});
</script>

{#if arcData}
	<div class="space-y-6 p-6">
		<!-- Arc header -->
		<div class="bg-card border border-border rounded-lg p-6">
			<div class="flex justify-between">
				<div class="space-y-3">
					<div class="flex items-center gap-3">
						<h2 class="text-2xl font-bold">Arc {arcData.arc}: {arcData.title}</h2>
						<span class="px-2 py-0.5 bg-primary/20 text-primary rounded text-sm">
							{arcData.status}
						</span>
					</div>

					<div class="text-sm text-muted-foreground flex gap-3 flex-wrap">
						<span>{arcData.episodeCount} Episodes</span>
						<span>•</span>
						<span>Audio: {arcData.audioLanguages}</span>
						<span>•</span>
						<span>Subtitles: {arcData.subtitleLanguages}</span>
						<span>•</span>
						<span>{arcData.resolution}</span>
					</div>
				</div>
			</div>
		</div>

		<!-- Episode list -->
		<div class="space-y-3">
			{#each arcData.episodes as episode}
				<EpisodeCard
					{episode}
					on:select={(e) => (selectedEpisode = e.detail)}
				/>
			{/each}
		</div>
	</div>
{/if}

<!-- MODAL -->
<EpisodeDetailsModal
	open={selectedEpisode !== null}
	onOpenChange={(open) => !open && (selectedEpisode = null)}
	episode={selectedEpisode}
/>
