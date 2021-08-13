<template>
    <div>
        <div v-for="folder in folders" :key="folder.id">
            {{folder.name}}
            <div v-for="bookmark in folder.bookmarks" :key="bookmark.id">
                <bookmark-component :bookmark="bookmark" />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { useApi } from '@/api/api'
import { defineComponent, ref } from 'vue'
import BookmarkComponent from './Bookmark.vue';

export default defineComponent({
    name: "FolderList",
    components: {
        BookmarkComponent
    },
    async setup() {
        const api = useApi();
        const response = await api.getFolders("bookmarks");
        const folders = ref(response.data);

        return {
            folders
        }
    },
})
</script>
