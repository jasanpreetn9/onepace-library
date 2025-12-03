<script lang="ts">
	import { arcs, selectedArc } from '$lib';
	import { api } from '$lib/api';

	let arcData: UnifiedArc | null = null;

	// Recalculate whenever store changes
	$: arcData = $arcs.find((a) => a.arc === $selectedArc) ?? null;

	async function toggleMonitor(ep: UnifiedEpisode) {
		ep.monitored = !ep.monitored;
		await api.toggleMonitor(ep.arc, ep.episode, ep.monitored);
	}

	async function download(ep: UnifiedEpisode) {
		await api.downloadEpisode(ep.crc32);
		ep.downloadStatus = 'queued';
	}
</script>

{#if arcData}
	<h2 style="margin-bottom: 1rem;">
		Arc {arcData.arc}: {arcData.title},
		<span style="opacity: 0.7; font-weight: 400;">
			{arcData.episodes.length} Episodes
		</span>
		<span style="opacity: 0.7; font-weight: 400;">
			| Audio Languages: {arcData.audioLanguages}</span
		>
		<span style="opacity: 0.7; font-weight: 400;">
			| Subtitle Languages: {arcData.subtitleLanguages}</span
		>
		<span style="opacity: 0.7; font-weight: 400;"> | Resolution: {arcData.resolution}</span>
		<span style="opacity: 0.7; font-weight: 400;"> | Status: {arcData.status}</span>
	</h2>

	{#each arcData.episodes as ep}
		<div class="episode">
			<div class="info">
				<h3>Episode {ep.episode} – {ep.title}</h3>
				{#if ep.filePath}
					<p>{ep.filePath}</p>
				{/if}
				<p style="opacity: 0.7;">{ep.description}</p>
				<div class="badges">
					<div style="margin-top: 0.4rem;">
					<span class="badge {ep.version}">
						{ep.version}
					</span>
				</div>
				<div style="margin-top: 0.4rem;">
					<span class="badge {ep.downloadStatus}">
						{ep.downloadStatus}
					</span>
				</div>
				</div>
			</div>

			<div>
				<button class="small" on:click={() => toggleMonitor(ep)}>
					{ep.monitored ? 'Unmonitor' : 'Monitor'}
				</button>

				{#if ep.downloadStatus === 'missing'}
					<button class="small download" on:click={() => download(ep)}>Download</button>
				{/if}
			</div>
		</div>
	{/each}
{/if}

<style>
	/* HEADER */
	h2 {
		font-size: 1.4rem;
		font-weight: 600;
		color: #ffffff;
		margin-bottom: 1.4rem;
		line-height: 1.6rem;
	}

	/* Inline metadata after arc title */
	h2 span {
		font-size: 0.9rem;
		color: #9ca0a8;
		margin-left: 0.2rem;
	}

	/* EPISODE CARD LAYOUT */
	.episode {
		background: #1d2027;
		border: 1px solid #2a2e35;
		padding: 1.2rem 1.4rem;
		border-radius: 10px;
		margin-bottom: 1rem;

		display: flex;
		justify-content: space-between;
		gap: 2rem;

		transition:
			background 0.2s ease,
			border-color 0.2s ease;
	}

	.episode:hover {
		background: #22262e;
		border-color: #3b414c;
	}

	/* INFO SECTION (LEFT) */
	.episode .info {
		flex: 1;
	}

	.episode .info h3 {
		margin-bottom: 0.4rem;
		font-size: 1.15rem;
		font-weight: 600;
		color: #e6e8ec;
	}

	.episode .info p {
		color: #a4a8b2;
		line-height: 1.5rem;
		font-size: 0.92rem;
		margin-bottom: 0.25rem;
	}

	/* File path */
	.episode .info p.filepath {
		font-family: monospace;
		font-size: 0.85rem;
		color: #7f8a99;
		background: #15171c;
		padding: 4px 6px;
		border-radius: 4px;
		width: fit-content;
	}
.badges {
	display: flex;
	gap: 1rem;
}
	/* BADGE */
	.badge {
		display: inline-block;
		padding: 5px 10px;
		border-radius: 6px;
		font-size: 0.78rem;
		font-weight: 600;
		color: white;
		text-transform: capitalize;

		margin-top: 0.5rem;
	}
	.badge.normal {
		background: #2d7b36;
	}
	.badge.extended {
		background: #295d9a;
	}
	.badge.imported {
		background: #2d7b36;
	}

	.badge.missing {
		background: #9b3a3a;
	}

	.badge.upgradable {
		background: #b08a1a;
	}

	.badge.queued {
		background: #295d9a;
	}

	.badge.downloading {
		background: #2372d8;
	}

	.badge.failed {
		background: #b43737;
	}

	/* RIGHT BUTTONS AREA */
	.episode > div:last-child {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		gap: 0.4rem;
		white-space: nowrap;
	}

	/* ACTION BUTTONS */
	button.small {
		padding: 0.4rem 1rem;
		background: #3b3f4a;
		border: none;
		color: #e5e7eb;
		border-radius: 5px;
		cursor: pointer;
		font-size: 0.85rem;
		font-weight: 500;

		transition:
			background 0.15s ease,
			transform 0.1s ease;
	}

	button.small:hover {
		background: #505562;
	}

	button.small:active {
		transform: scale(0.97);
		background: #5d626f;
	}

	/* BLUE DOWNLOAD BUTTON */
	button.small.download {
		background: #2d89ef;
		color: white;
	}

	button.small.download:hover {
		background: #2073c7;
	}
</style>
