<template>
  <div class="flex justify-center flex-wrap">
    <suspense>
      <template #fallback>Loading...</template>
      <template #default>
        <bookmarks-list />
      </template>
    </suspense>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BookmarksList from "@/components/BookmarksList.vue";

export default defineComponent({
  name: "Bookmarks",
  components: {
    BookmarksList,
  },
});
</script>
