<template>
  <router-link
    v-if="routerLink"
    :to="linkTo"
    :class="classes"
    class="transition duration-300 p-2 shadow"
  >
    <slot></slot>
  </router-link>
  <button v-else :class="classes" class="p-2 shadow transition duration-300">
    <slot></slot>
  </button>
</template>

<script lang="ts">
import { computed, defineComponent } from "vue";

export default defineComponent({
  name: "Button",
  props: {
    rounded: {
      type: Boolean,
      default: false,
    },
    type: {
      type: String,
      default: "blue",
    },
    dark: {
      type: Boolean,
      default: false,
    },
    hoverColor: {
      type: String,
      default: "blue-900",
    },
    routerLink: {
      type: Boolean,
      default: false,
    },
    linkTo: {
      type: String,
      default: "#",
    },
  },
  emits: ["update:dark"],
  setup(props) {
    const classes = computed(() => {
      return {
        rounded: props.rounded,
        [`bg-${props.type}-700`]: props.dark,
        "text-white": props.dark,
        [`hover:bg-${props.type}-800`]: props.dark,
        [`bg-${props.type}-200`]: !props.dark,
        [`text-${props.type}-700`]: !props.dark,
        [`hover:bg-${props.type}-300`]: !props.dark,
      };
    });
    return { classes };
  },
});
</script>
