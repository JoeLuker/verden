declare global {
	namespace App {
		// Uncomment or add as needed:
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}
}

/// <reference types="svelte" />

// Declare ImportMetaEnv to hold environment variables
interface ImportMetaEnv {
	VITE_API_URL: string;
	// Add other environment variables as needed
}

// Extend the ImportMeta interface to include env
interface ImportMeta {
	readonly env: ImportMetaEnv;
}

// Existing declaration for importing SVGs
declare module '*.svg' {
	const content: string;
	export default content;
}

export {};
