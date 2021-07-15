<template>
  <div>
    <ul>
      <li>
        <input id="theme-light" v-model="themeSelection" type="radio" value="light" />
        <label class="p-4" for="theme-light">Light</label>
      </li>
      <li>
        <input id="theme-dark" v-model="themeSelection" type="radio" value="dark" />
        <label class="p-4" for="theme-dark">Dark</label>
      </li>
      <li>
        <input id="theme-os" v-model="themeSelection" type="radio" value="os" />
        <label class="p-4" for="theme-os">OS Specified</label>
      </li>
    </ul>
    <dm-button :dark="state.isDarkmode()" type="primary" rounded @click="handleThemeSave()">Save</dm-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useState } from "../store/store";
import DmButton from "../components/Button.vue";
import router from "@/router";

export default defineComponent({
  name: "UserSettings",
  components: {
    DmButton,
  },
  setup() {
    const state = useState();
    const themeSelection = ref("");
    const handleThemeSave = () => {
      if (themeSelection.value === "light") {
        state.setDarkmode(false);
      } else if (themeSelection.value === "dark") {
        state.setDarkmode(true);
      } else if (themeSelection.value === "os") {
        state.darkmodeOSPreference();
      } else {
        console.error("invalid theme value");
      }
      router.go(0);
    };
    return {
      themeSelection,
      state,
      handleThemeSave,
    };
  },
});
</script>
