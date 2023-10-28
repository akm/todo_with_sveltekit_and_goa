import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import vitePluginRequire from 'vite-plugin-require';

export default defineConfig({
	plugins: [vitePluginRequire.default(), sveltekit()]
});
