import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import reactRefresh from '@vitejs/plugin-react-refresh'

export default defineConfig({
	plugins: [sveltekit(), reactRefresh()],
	server: {
	  hmr: {
		clientPort: process.env.VITE_CLIENT_PORT || null
	  }
	}
});
