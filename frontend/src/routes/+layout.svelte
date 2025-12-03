<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';
	import { arcs, selectedArc } from '$lib';
	import "$lib/app.css";

	// Load arcs when the app starts
	onMount(async () => {
		const list = await api.getAllEpisodes();
		arcs.set(list);

		if (list.length > 0) {
			selectedArc.set(list[0].arc);
		}
	});

	async function scanLibrary() {
		await api.scanLibrary();
		const list = await api.getAllEpisodes();
		arcs.set(list);
	}

	async function scanDownloads() {
		await api.scanDownloads();
		const list = await api.getAllEpisodes();
		arcs.set(list);
	}
</script>

<div class="sidebar">
	<h2>Arcs</h2>

	{#each $arcs as arc}
		<button 
			class:selected={$selectedArc === arc.arc}
			on:click={() => selectedArc.set(arc.arc)}
		>
			{arc.arc} – {arc.title}
		</button>
	{/each}
</div>

<div class="main">
	<div class="header">
		<h1>One Pace Library</h1>

		<div>
			<button on:click={scanLibrary}>Scan Library</button>
			<button on:click={scanDownloads}>Scan Downloads</button>
		</div>
	</div>
	
	<div class="content">
		<slot />
	</div>
</div>

<style>
/* LAYOUT ROOT */
.sidebar {
	width: 300px;
	background: #13151a;
	border-right: 1px solid #2b2f36;
	padding: 1.2rem 1rem;
	overflow-y: auto;
	display: flex;
	flex-direction: column;
	gap: 0.4rem;
}

/* Sidebar title */
.sidebar h2 {
	font-size: 0.95rem;
	font-weight: 600;
	text-transform: uppercase;
	letter-spacing: 1px;
	color: #7c818c;
	margin-bottom: 1rem;
	padding-left: 2px;
}

/* Sidebar buttons */
.sidebar button {
	background: transparent;
	border: none;
	padding: 0.5rem 0.7rem;
	text-align: left;
	border-radius: 6px;

	font-size: 0.95rem;
	color: #d0d3d9;
	cursor: pointer;

	transition: background 0.15s ease, color 0.15s ease;
}

/* Hover + selected styles */
.sidebar button:hover {
	background: #20242c;
	color: #ffffff;
}

.sidebar button.selected {
	background: #2c323d;
	color: #ffffff;
	font-weight: 500;
}

/* Subtle left highlight bar */
.sidebar button.selected {
	position: relative;
}

.sidebar button.selected::before {
	content: "";
	position: absolute;
	left: -12px;
	top: 10%;
	height: 80%;
	width: 4px;
	background: #2d89ef;
	border-radius: 2px;
}

/* Custom scrollbar */
.sidebar::-webkit-scrollbar {
	width: 8px;
}

.sidebar::-webkit-scrollbar-thumb {
	background: #2c323d;
	border-radius: 4px;
}

.sidebar::-webkit-scrollbar-thumb:hover {
	background: #3d4450;
}

/* HEADER */
.header {
	height: 56px;
	background: #181b22;
	border-bottom: 1px solid #2b2f36;

	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: 0 1rem;
}

/* Title */
.header h1 {
	font-size: 1.2rem;
	font-weight: 600;
	color: #e3e5e8;
}

/* Action buttons */
.header button {
	background: #2d89ef;
	border: none;
	padding: 0.35rem 0.9rem;
	border-radius: 5px;
	color: white;
	cursor: pointer;
	font-size: 0.9rem;
}

.header button:hover {
	background: #2276d2;
}

/* MAIN */
.main {
	flex-grow: 1;
	overflow-y: auto;
	background: #1b1e24;
}

.content {
	padding: 1.5rem;
}

/* Smooth scroll */
html {
	scroll-behavior: smooth;
}
</style>
