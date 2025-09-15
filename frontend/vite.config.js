import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
    plugins: [vue()],
    server: {
        port: 3000,
        host: '0.0.0.0',
        proxy: {
            '/api': {
                target: 'http://localhost:8080',
                changeOrigin: true
                // 注意：删除了 rewrite 规则，保持完整的 /api 路径
            }
        }
    },
    build: {
        outDir: 'dist'
    }
})