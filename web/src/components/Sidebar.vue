<template>
  <aside v-if="mountSidebar">
    <transition name="slide" appear @after-leave="close()">
      <div v-show="closeSidebar" class="h-full shadow-2xl" :class="classes">
        <slot></slot>
      </div>
    </transition>
  </aside>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from "vue";

export default defineComponent({
  name: "Sidebar",
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
    modelValue: {
      type: Boolean,
      default: true,
    },
  },
  emits: ["update:modelValue", "update:dark"],
  setup(props, { emit }) {
    const classes = computed(() => {
      return {
        [`bg-gray-100`]: !props.dark && !props.bgColor,
        ["bg-gray-800"]: props.dark && !props.bgColor,
        [`bg-${props.bgColor}`]: !!props.bgColor,
        [`text-${props.color}-700`]: !props.dark && !props.textColor,
        [`text-gray-100`]: props.dark && !props.textColor,
        [`text-${props.textColor}`]: !!props.textColor,
      };
    });

    const mountSidebar = ref(props.modelValue);
    const closeSidebar = ref(props.modelValue);

    const close = () => {
      emit("update:modelValue", false);
      mountSidebar.value = props.modelValue;
    };

    watch(
      () => props.modelValue,
      (value) => {
        if (value) mountSidebar.value = true;
        closeSidebar.value = value;
      }
    );

    return {
      classes,
      close,
      mountSidebar,
      closeSidebar,
    };
  },
});
</script>

<style>
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.2s ease;
}
.slide-enter-from,
.slide-leave-to {
  transform: translateX(-100%);
  transition: all 150ms ease-in 0s;
}
</style>

