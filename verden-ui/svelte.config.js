import adapter from '@sveltejs/adapter-auto';
// import { vitePreprocess } from '@sveltejs/kit/vite';
import sveltePreprocess from 'svelte-preprocess';

const config = {
    preprocess: sveltePreprocess({
        // options for svelte-preprocess
        typescript: {
            // TypeScript options
        },
    }),

    kit: {
        adapter: adapter(),
        // other kit options...
    },
};

export default config;
