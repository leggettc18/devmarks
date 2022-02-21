<template>
  <div class="flex justify-center flex-wrap">
    <suspense>
      <template #fallback>Loading...</template>
      <template #default>
        <div>
          <div v-show="view === ViewType.Folders">
            <folder-list />
          </div>
          <div v-show="view === ViewType.Bookmarks">
            <bookmarks-list />
          </div>
        </div>
      </template>
    </suspense>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import BookmarksList from "@/components/BookmarksList.vue";
import FolderList from "@/components/FolderList.vue";

export default defineComponent({
  name: "Bookmarks",
  components: {
    BookmarksList,
    FolderList,
  },
  setup() {
    enum ViewType {
      Folders = 1,
      Bookmarks,
      Tags,
    }
    const view = ref(ViewType.Bookmarks);

    return {
      view,
      ViewType,
    };
  },
});
</script>
