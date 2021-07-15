<template>
  <div class="p-4 flex justify-around cursor-pointer" :class="classes">
    <slot></slot>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent } from "vue";

export default defineComponent({
  name: "Card",
  props: {
    color: {
      type: String,
      default: "indigo",
    },
    bgColor: {
      type: String,
      default: "",
    },
    textColor: {
      type: String,
      default: "",
    },
    dark: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["update:dark"],
  setup(props) {
    const classes = computed(() => {
      return {
        [`bg-gray-800`]: props.dark && !props.bgColor,
        [`bg-${props.bgColor}`]: !!props.bgColor,
        [`bg-${props.color}-200`]: !props.dark && !props.bgColor,
        [`text-gray-200`]: props.dark && !props.textColor,
        [`text-${props.color}-700`]: !props.dark && !props.textColor,
        [`text-${props.textColor}`]: !!props.textColor,
      };
    });

    return {
      classes,
    };
  },
});
</script>
