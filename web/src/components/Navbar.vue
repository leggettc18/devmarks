<template>
  <header class="w-full" :class="classes">
    <nav class="w-full items-center p-2">
      <div class="flex justify-between items-center">
        <slot name="nav-left"></slot>
        <slot name="nav-right"></slot>
      </div>
    </nav>
  </header>
</template>

<script lang="ts">
import { computed, defineComponent, toRefs } from "vue";
export default defineComponent({
  name: "NavBar",
  props: {
    dark: {
      type: Boolean,
      default: false,
    },
    color: {
      type: String,
      default: "blue",
    },
    bgColor: {
      type: String,
      default: "",
    },
    textColor: {
      type: String,
      default: "",
    },
  },
  emits: ["update:dark"],
  setup(props) {
    const { dark, color, bgColor, textColor } = toRefs(props);
    const classes = computed(() => {
      return {
        [`bg-${color.value}-700`]: !dark.value && !bgColor.value,
        [`bg-gray-800`]: dark.value && !bgColor.value,
        [`bg-${bgColor.value}`]: !!bgColor.value,
        ["text-white"]: !textColor.value,
        [`text-${textColor.value}`]: !!textColor.value,
      };
    });

    return {
      classes,
    };
  },
});
</script>