<template>
  <div>
    <h1 class="text-2xl text-center">Register</h1>
    <dm-input
      v-model="form.email"
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
      v-model="form.password"
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
        @click.prevent="register(form)"
      >Login</dm-button>
      <dm-button type="danger" :dark="state.isDarkmode()" rounded router-link link-to="/">Cancel</dm-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useState } from "../store/store";
import { Credentials } from "@/models/auth";
import { useRegisterMutation } from "@/generated/graphql";
import router from "@/router";
import DmButton from "@/components/Button.vue";
import DmInput from "@/components/Input.vue";
import { GraphQLError } from "graphql";

export default defineComponent({
  name: "Register",
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

    const registerErrors = ref({
      email: null,
      password: null,
    } as {
      email: null | GraphQLError[];
      password: null | GraphQLError[];
    });

    const {
      mutate: register,
      onDone,
      error: registerError,
      loading,
    } = useRegisterMutation(() => ({
      variables: credentials.value,
      errorPolicy: "all",
    }));

    onDone((register) => {
      if (register.data?.register?.token && register.data.register.user) {
        state.storeToken({ token: register.data?.register.token });
        state.storeUser(register.data?.register?.user);
        router.push("/home");
      }
      if (register.errors) {
        registerErrors.value.email = register.errors.filter((e) => {
          return e.extensions?.field === "email";
        });
        registerErrors.value.password = register.errors.filter((e) => {
          return e.extensions?.field === "password";
        });
      }
    });
    return {
      state,
      register,
      credentials,
      registerError,
      registerErrors,
      loading,
    };
  },
  data() {
    return {
      form: this.credentials,
    };
  },
});
</script>
