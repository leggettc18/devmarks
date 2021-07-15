<template>
  <div class="flex items-center">
    <div>
      <dm-input
        v-model="colorSelected"
        type="text"
        name="color"
        color="primary"
        label="Pick a color"
        @input="$emit('update:modelValue', $event.target.value)"
      ></dm-input>
    </div>
    <div class="relative ml-3 mt-8">
      <button
        type="button"
        class="w-10 h-10 rounded-full focus:outline-none focus:shadow-outline inline-flex p-2 shadow"
        :style="`background: ${colorSelected}; color: white`"
        @click="isOpen = !isOpen"
      >
        <svg class="w-6 h-6 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <path
            fill="none"
            d="M15.584 10.001L13.998 8.417 5.903 16.512 5.374 18.626 7.488 18.097z"
          />
          <path
            d="M4.03,15.758l-1,4c-0.086,0.341,0.015,0.701,0.263,0.949C3.482,20.896,3.738,21,4,21c0.081,0,0.162-0.01,0.242-0.03l4-1 c0.176-0.044,0.337-0.135,0.465-0.263l8.292-8.292l1.294,1.292l1.414-1.414l-1.294-1.292L21,7.414 c0.378-0.378,0.586-0.88,0.586-1.414S21.378,4.964,21,4.586L19.414,3c-0.756-0.756-2.072-0.756-2.828,0l-2.589,2.589l-1.298-1.296 l-1.414,1.414l1.298,1.296l-8.29,8.29C4.165,15.421,4.074,15.582,4.03,15.758z M5.903,16.512l8.095-8.095l1.586,1.584 l-8.096,8.096l-2.114,0.529L5.903,16.512z"
          />
        </svg>
      </button>

      <TransitionRoot
        as="template"
        enter="transition ease-out duration-100 transform"
        enter-from="opacity-0 scale-95"
        enter-to="opacity-100 scale-100"
        leave="transition ease-in duration-75 transform"
        leave-from="opacity-100 scale-100"
        leave-to="opacity-100 scale-95"
      >
        <div
          v-show="isOpen"
          ref="colorPopup"
          class="origin-top-left absolute right-0 -top-40 mt-2 w-40 rounded-md shadow-lg"
        >
          <div class="rounded-md bg-white shadow-xs px-4 py-3">
            <div class="flex flex-wrap -mx-2">
              <div v-for="(color, index) in colors" :key="index">
                <div class="px-2">
                  <template v-if="colorSelected === color">
                    <div
                      class="w-8 h-8 inline-flex rounded-full cursor-pointer border-4 border-white"
                      :style="`background: ${color}; box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.2);`"
                    ></div>
                  </template>

                  <template v-if="colorSelected != color">
                    <div
                      role="checkbox"
                      tabindex="0"
                      :aria-checked="colorSelected"
                      class="w-8 h-8 inline-flex rounded-full cursor-pointer border-4 border-white focus:outline-none focus:shadow-outline"
                      :style="`background: ${color};`"
                      @click="selectColor(color)"
                      @keydown.enter="selectColor(color)"
                    ></div>
                  </template>
                </div>
              </div>
            </div>
          </div>
        </div>
      </TransitionRoot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watchEffect } from "vue";
import { TransitionRoot } from "@headlessui/vue";
import DmInput from "@/components/Input.vue";
import { onClickOutside } from "@vueuse/core";

export default defineComponent({
  name: "ColorPicker",
  components: {
    TransitionRoot,
    DmInput,
  },
  props: {
    modelValue: {
      type: String,
      default: "#2196f3",
    },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const colorPopup = ref(null);
    const isOpen = ref(false);
    const colors = [
      "#2196F3",
      "#009688",
      "#9C27B0",
      "#FFEB3B",
      "#afbbc9",
      "#4CAF50",
      "#2d3748",
      "#f56565",
      "#ed64a6",
    ];
    const colorSelected = ref("#2196f3");
    const selectColor = (newColor: string) => {
      colorSelected.value = newColor;
      emit("update:modelValue", newColor);
      isOpen.value = false;
    };
    const inputValue = ref(props.modelValue);

    onClickOutside(colorPopup, () => {
      isOpen.value = false;
    });

    watchEffect(() => {
      colorSelected.value = inputValue.value;
    });

    return {
      isOpen,
      colors,
      colorSelected,
      selectColor,
      colorPopup,
    };
  },
});
</script>
