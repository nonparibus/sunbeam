import { defineConfig, PluginOption } from "vite";
import react from "@vitejs/plugin-react";

const rewriteSlashToIndexHtml = () => {
  return {
    name: 'rewrite-slash-to-index-html',
    apply: 'serve' as "serve",
    enforce: 'post' as "post",
    configureServer(server) {
      // rewrite / as index.html
      server.middlewares.use('/', (req, _, next) => {
        if (req.url === '/') {
          req.url = '/index.html'
        }
        next()
      })
    },
  }
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), rewriteSlashToIndexHtml()],
  publicDir: "./public",
  appType: "mpa",
  build: {
    sourcemap: true,
  },
});
