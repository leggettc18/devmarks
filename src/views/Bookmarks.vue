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
    <div class="flex justify-center flex-wrap">
      <template v-if="loading">Loading...</template>
      <template v-else-if="error">Error: {{error.message}}</template>
      <template v-else-if="bookmarks">
        <div v-for="(bookmark, i) in bookmarks" :key="i">
          <card :dark="state.isDarkmode()" color="primary" class="m-4 rounded-lg shadow">
            <div class="flex space-x-4 items-center">
              <a :href="'https://' + bookmark.url" class="hover:underline">{{bookmark.name}}</a>
              <dm-button
                type="info"
                rounded
                :dark="state.isDarkmode()"
                @click="handleEdit(bookmark)"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
                  />
                  <path
                    fill-rule="evenodd"
                    d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
                    clip-rule="evenodd"
                  />
                </svg>
              </dm-button>
              <dm-button
                type="danger"
                rounded
                :dark="state.isDarkmode()"
                @click="handleDelete(bookmark.id)"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                  />
                </svg>
              </dm-button>
            </div>
          </card>
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, Ref, ref } from "vue";
import { useState } from "@/store/store";
import { Bookmark, BookmarkCreate, BookmarkUpdate } from "@/models/bookmark";
import { useResult } from "@vue/apollo-composable";
import {
  GetBookmarksDocument,
  GetBookmarksQuery,
  useGetBookmarksQuery,
  useNewBookmarkMutation,
  useDeleteBookmarkMutation,
  useUpdateBookmarkMutation,
} from "../generated/graphql";
import {
  TransitionChild,
  TransitionRoot,
  Dialog,
  DialogOverlay,
  DialogTitle,
} from "@headlessui/vue";
import DmButton from "@/components/Button.vue";
import DmInput from "@/components/Input.vue";
import Card from "@/components/Card.vue";
import ColorPicker from "@/components/ColorPicker.vue";

export default defineComponent({
  name: "Bookmarks",
  components: {
    DmButton,
    DmInput,
    Card,
    TransitionChild,
    TransitionRoot,
    Dialog,
    DialogOverlay,
    DialogTitle,
    ColorPicker,
  },
  setup() {
    const editing = ref(false);
    const state = useState();
    const token = state.getToken();
    const dialogVisible: Ref<boolean> = ref(false);
    if (!token) {
      return {};
    }
    const openDialog = () => {
      dialogVisible.value = true;
    };
    const closeDialog = () => {
      dialogVisible.value = false;
    };

    const { result, loading, error } = useGetBookmarksQuery();
    const bookmarks = useResult(result);

    const dialogBookmark = ref({
      id: "",
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
      dialogBookmark.value = { id: "", ...newBookmark.value };
      openDialog();
    };

    const handleClose = () => {
      editing.value = false;
      closeDialog();
    };

    const { mutate: submitNewBookmark } = useNewBookmarkMutation(() => ({
      variables: newBookmark.value,
      update: (cache, { data: submitBookmark }) => {
        const data = cache.readQuery<GetBookmarksQuery>({
          query: GetBookmarksDocument,
        });
        const submittedBookmark = {
          name: submitBookmark?.newBookmark.name,
          url: submitBookmark?.newBookmark.url,
          color: submitBookmark?.newBookmark.color,
        } as Bookmark;
        if (data?.bookmarks) {
          cache.writeQuery<GetBookmarksQuery>({
            query: GetBookmarksDocument,
            data: { bookmarks: [...data.bookmarks, submittedBookmark] },
          });
        }
      },
    }));

    const updatedBookmark = ref({
      id: "",
      name: "",
      url: "",
      color: "",
    } as BookmarkUpdate);

    const handleEdit = (bookmark: Bookmark) => {
      dialogBookmark.value = { ...bookmark };
      dialogVisible.value = true;
      editing.value = true;
    };

    const { mutate: updateBookmark } = useUpdateBookmarkMutation(() => ({
      variables: updatedBookmark.value,
      update: (cache, { data: updateBookmark }) => {
        const data = cache.readQuery<GetBookmarksQuery>({
          query: GetBookmarksDocument,
        });
        const bookmark = {
          id: updateBookmark?.updateBookmark.id,
          name: updateBookmark?.updateBookmark.name,
          url: updateBookmark?.updateBookmark.url,
          color: updateBookmark?.updateBookmark.color,
        } as Bookmark;
        if (data?.bookmarks) {
          const indexToUpdate = data.bookmarks.findIndex(
            (oldBookmark) => oldBookmark.id === bookmark.id
          );
          const updatedBookmarks = [...data.bookmarks];
          updatedBookmarks[indexToUpdate] = bookmark;
          cache.writeQuery<GetBookmarksQuery>({
            query: GetBookmarksDocument,
            data: { bookmarks: updatedBookmarks },
          });
        }
      },
    }));

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

    const deleteId = ref<number | null>(null);

    const { mutate: deleteBookmark } = useDeleteBookmarkMutation(() => ({
      variables: { id: `${deleteId.value}` },
      update: (cache) => {
        const data = cache.readQuery<GetBookmarksQuery>({
          query: GetBookmarksDocument,
        });
        if (data?.bookmarks) {
          const bookmarks = data.bookmarks.filter((bookmark) => {
            console.log(bookmark.id !== `${deleteId.value}`);
            return bookmark.id !== `${deleteId.value}`;
          });
          console.log(bookmarks);
          cache.writeQuery<GetBookmarksQuery>({
            query: GetBookmarksDocument,
            data: { bookmarks },
          });
        }
      },
    }));

    const handleDelete = (id: number) => {
      deleteId.value = id;
      deleteBookmark();
    };

    return {
      editing,
      state,
      dialogVisible,
      openDialog,
      closeDialog,
      bookmarks,
      loading,
      error,
      dialogBookmark,
      handleNew,
      handleClose,
      handleSubmit,
      handleDelete,
      handleEdit,
    };
  },
});
</script>
