import { createApp } from "vue";
import { stateSymbol, createState } from "./store/store";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import "@/assets/tailwind.css";
import { createApi, DefaultDevmarksClient } from "./api/api";
import { Configuration } from "./api/client";

const config = new Configuration({
  accessToken: () => {
    const token = localStorage.getItem("user-token");
    if (token === null) {
      return "";
    }
    return token;
  }
});

const devmarksClient = createApi(config);

createApp(App)
  .provide(DefaultDevmarksClient, devmarksClient)
  .provide(stateSymbol, createState())
  .use(router)
  .mount("#app");
