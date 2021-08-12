<template>
  <div>
    <h1 class="text-2xl text-center">Login</h1>
    <dm-input
      v-model="form.email"
      type="email"
      name="email"
      label="E-Mail"
      color="primary"
      :error="loginErrors.emaili != null"
      @update:modelValue="loginErrors.email = null"
    ></dm-input>
    <template v-if="loginErrors.email">
      <div v-for="(error, i) in loginErrors.email" :key="i">
        <span v-if="error.extensions" class="text-danger-700">{{error.extensions.message}}</span>
      </div>
    </template>
    <dm-input
      v-model="form.password"
      type="password"
      name="password"
      label="Password"
      color="primary"
      :error="loginErrors.password != null"
      @update:modelValue="loginErrors.password = null"
    ></dm-input>
    <template v-if="loginErrors.password">
      <div v-for="(error, i) in loginErrors.password" :key="i">
        <span v-if="error.extensions" class="text-danger-700">{{error.extensions.message}}</span>
      </div>
    </template>
    <div class="flex justify-center space-x-4">
      <dm-button type="primary" :dark="state.isDarkmode()" rounded @click.prevent="handleLogin()">Login</dm-button>
      <dm-button type="danger" :dark="state.isDarkmode()" rounded router-link link-to="/">Cancel</dm-button>
    </div>
    <div v-if="loginErrors">
      <div v-for="(e, i) of loginErrors.email" :key="i">{{e.extensions.message}}</div>
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
  name: "Login",
  components: {
    DmButton,
    DmInput,
  },
  setup() {
    const state = useState();

    const credentials = ref({
      email: "",
      password: "",
    } as Credentials);

    const loginErrors = ref({
      email: null,
      password: null,
    } as {
      email: null;
      password: null;
    });

    const api = useApi();

    const handleLogin = async () => {
      const response = await api.userApi.login(credentials.value);
      state.storeToken({ token: response.data.token})
      state.storeUser(await (await api.userApi.getUser()).data);
      router.push("/home");
    };

    return {
      state,
      loginErrors,
      credentials,
      handleLogin,
    };
  },
  data() {
    return {
      form: this.credentials,
    };
  },
});
</script>
