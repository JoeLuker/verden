// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}
}


/// <reference types="svelte" />
interface ImportMetaEnv {
    VITE_API_URL: string;
    // add other environment variables as needed
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}


declare module '*.svg' {
    const content: any;
    export default content;
}

export {};
