<template>
  <div>
    <h1 class="text-2xl text-center">Register</h1>
    <dm-input
      v-model="credentials.email"
      type="email"
      label="E-Mail"
      name="email"
      color="primary"
      :error="registerErrors.email != null"
      @update:modelValue="registerErrors.email = null"
    ></dm-input>
    <template v-if="registerErrors.email">
      <div v-for="(error, i) in registerErrors.email" :key="i">
        <span v-if="error.extensions" class="text-danger">{{error.extensions.message}}</span>
      </div>
    </template>
    <dm-input
      v-model="credentials.password"
      type="password"
      label="Password"
      name="password"
      color="primary"
      :error="registerErrors.password != null"
      @update:modelValue="registerErrors.password = null"
    ></dm-input>
    <template v-if="registerErrors.password">
      <div v-for="(error, i) in registerErrors.password" :key="i">
        <span v-if="error.extensions" class="text-danger">{{error.extensions.message}}</span>
      </div>
    </template>
    <div class="flex justify-center space-x-4">
      <dm-button
        type="primary"
        :dark="state.isDarkmode()"
        rounded
        @click.prevent="handleRegister"
      >Register</dm-button>
      <dm-button type="danger" :dark="state.isDarkmode()" rounded router-link link-to="/">Cancel</dm-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useState } from "../store/store";
import { Credentials } from "@/models/auth";
import router from "@/router";
import DmButton from "@/components/Button.vue";
import DmInput from "@/components/Input.vue";
import { useApi } from "@/api/api";

export default defineComponent({
  name: "Register",
  components: {
    DmButton,
    DmInput,
  },
  setup() {
    const state = useState();
    const api = useApi();

    const credentials = ref({
      email: "",
      password: "",
    } as Credentials);

    const registerErrors = ref({
      email: null,
      password: null,
    } as {
      email: null;
      password: null;
    });

    const handleRegister = async () => {
      await api.userApi.register(credentials.value);
      router.push("/login");
    };

    
    return {
      state,
      credentials,
      registerErrors,
      handleRegister
    };
  },
  data() {
    return {
      form: this.credentials,
    };
  },
});
</script>
