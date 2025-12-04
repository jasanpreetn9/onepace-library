// place files you want to import through the `$lib` alias in this folder.
import { writable } from "svelte/store";

export const arcs = writable<UnifiedArc[]>([]);
export const selectedArc = writable<number | null>(null);
export	const sidebarOpen = writable(false);
