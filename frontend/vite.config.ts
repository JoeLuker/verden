import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	server: {
		host: '0.0.0.0',

		hmr: {
		  protocol: 'ws',
		  clientPort: 5173,
		},
		// Uncomment the following line to use polling
		watch: {
		  usePolling: true,
		},
	  },
});
