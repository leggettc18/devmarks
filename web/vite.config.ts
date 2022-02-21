// vite.config.js

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
const path = require("path");
export default defineConfig({
    server: {
        host: "0.0.0.0"
    },
    plugins: [vue()],
    resolve: {
        alias: {
            "@": "/src",
        },
    },
})