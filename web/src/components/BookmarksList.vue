<template>
  <div>
    <div class="flex justify-center">
      <dm-button type="primary" :dark="state.isDarkmode()" rounded @click="handleNew()">Add Bookmark</dm-button>
    </div>
    <TransitionRoot appear :show="dialogVisible" as="template">
      <Dialog as="div" class="fixed inset-0 z-10 overflow-y-auto" @close="closeDialog()">
        <div class="flex justify-center items-center">
          <div class="min-h-screen px-4 text-center">
            <transition-child
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0"
              enter-to="opacity-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100"
              leave-to="opacity-0"
            >
              <dialog-overlay class="fixed inset-0 bg-black bg-opacity-30" />
            </transition-child>

            <span class="inline-block h-screen align-middle" aria-hidden="true">&#8203;</span>

            <transition-child
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <div
                class="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform shadow-xl rounded-2xl"
                :class="{['bg-gray-700']: state.isDarkmode(), ['text-gray-100']: state.isDarkmode(), ['bg-white']: !state.isDarkmode()}"
              >
                <dialog-title
                  as="h3"
                  class="text-lg font-medium leading-6 pb-4"
                  :class="{['text-gray-900']: !state.isDarkmode(), ['text-gray-200']: state.isDarkmode()}"
                >Add Bookmark</dialog-title>
                <div>
                  <dm-input
                    v-model="dialogBookmark.name"
                    type="text"
                    name="name"
                    label="Name"
                    color="primary"
                  ></dm-input>
                  <div class="flex items-baseline gap-4">
                    <p
                      class="rounded-lg p-1"
                      :class="{['bg-primary-300']: !state.isDarkmode(),  ['bg-gray-500']: state.isDarkmode()}"
                    >https://</p>
                    <dm-input
                      v-model="dialogBookmark.url"
                      type="text"
                      name="url"
                      label="URL"
                      color="primary"
                    ></dm-input>
                  </div>
                  <color-picker v-model="dialogBookmark.color" class="pb-8"></color-picker>
                </div>
                <span class="dialog-footer flex space-x-4">
                  <dm-button
                    :dark="state.isDarkmode()"
                    type="primary"
                    @click="handleSubmit()"
                  >Submit</dm-button>
                  <dm-button :dark="state.isDarkmode()" type="danger" @click="handleClose()">Cancel</dm-button>
                </span>
              </div>
            </transition-child>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
    <div v-for="(bookmark, i) in bookmarks" :key="i">
      <bookmark-component :bookmark="bookmark" @edit="handleEdit" @delete="handleDelete"/>
    </div>
  </div>
</template>

<script lang="ts">
import { useApi } from "@/api/api";
import { useState } from "@/store/store";
import { Bookmark } from "@/api/client/api";
import { defineComponent, Ref, ref } from "vue";
import BookmarkComponent from "@/components/Bookmark.vue";
import DmInput from "@/components/Input.vue";
import DmButton from "@/components/Button.vue";
import ColorPicker from "@/components/ColorPicker.vue"
import {
  TransitionChild,
  TransitionRoot,
  Dialog,
  DialogOverlay,
  DialogTitle
} from "@headlessui/vue"
import router from "@/router";
import { BookmarkCreate, BookmarkUpdate } from "@/models/bookmark";

export default defineComponent({
  name: "Bookmarks",
  components: {
    BookmarkComponent,
    TransitionChild,
    TransitionRoot,
    Dialog,
    DialogOverlay,
    DialogTitle,
    DmInput,
    DmButton,
    ColorPicker,
  },
  emits: ["edit", "delete"],
  async setup() {
    const state = useState();
    const bookmarks: Ref<Bookmark[]> = ref([]);
    const editing = ref(false);
    const token = state.getToken();
    const dialogVisible: Ref<boolean> = ref(false);
    const api = useApi();
    if (!token) {
      return {};
    }

    const response = await api.getBookmarks();
    if (!response.success) {
        if (response.statusCode === 403) {
          state.logOut();
          router.push('/login');
        }
    } else {
      if (response.data) bookmarks.value = response.data;
    }
    const openDialog = () => {
      dialogVisible.value = true;
    };
    const closeDialog = () => {
      dialogVisible.value = false;
    };

    const dialogBookmark = ref({
      id: 0,
      name: "",
      color: "",
      url: "",
    } as Bookmark);

    const newBookmark = ref({
      name: "",
      url: "",
      color: "",
    } as BookmarkCreate);

    const handleNew = () => {
      editing.value = false;
      dialogBookmark.value = { id: 0, ...newBookmark.value };
      openDialog();
    };

    const handleClose = () => {
      editing.value = false;
      closeDialog();
    };

    const submitNewBookmark = async () => {
      const response = await api.newBookmark(newBookmark.value);
      if (response.success && response.data) {
        bookmarks.value?.push(response.data)
      } else {
        console.error(response.message);
      }
    };
    
    const updatedBookmark = ref({
      id: 0,
      name: "",
      url: "",
      color: "",
    } as BookmarkUpdate);

    const handleEdit = (bookmark: Bookmark) => {
      dialogBookmark.value = { ...bookmark };
      dialogVisible.value = true;
      editing.value = true;
    };

    const updateBookmark = async () => {
      const response = await api.updateBookmark(updatedBookmark.value);
      if (response.success && response.data) {
        const index = bookmarks.value.findIndex(bookmark => bookmark.id === response.data?.id);
        console.log(index);
        if (index !== undefined) bookmarks.value.splice(index, 1, response.data);
      } else {
        console.error(response.message);
      }
    }

    const handleSubmit = () => {
      if (editing.value) {
        updatedBookmark.value = dialogBookmark.value;
        updateBookmark();
        editing.value = false;
      } else {
        newBookmark.value = dialogBookmark.value;
        submitNewBookmark();
      }
      dialogVisible.value = false;
    };

    const deleteId = ref<number>(0);

    const deleteBookmark = async () => {
      const response = await api.deleteBookmark(deleteId.value);
      if (response.success) {
        const index = bookmarks.value.findIndex(bookmark => bookmark.id === deleteId.value)
        bookmarks.value.splice(index, 1);
      }
    }

    const handleDelete = (bookmark: Bookmark) => {
      deleteId.value = bookmark.id;
      deleteBookmark();
    };

    return {
      state,
      bookmarks,
      handleEdit,
      handleClose,
      handleNew,
      handleSubmit,
      handleDelete,
      dialogBookmark,
      dialogVisible,
    };
  },
});
</script>
