import { createApp } from "vue";
import { stateSymbol, createState } from "./store/store";
import { DefaultApolloClient } from "@vue/apollo-composable";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import {
  ApolloClient,
  ApolloLink,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client/core";
import "@/assets/tailwind.css";
import { createApi, DefaultDevmarksClient, DevmarksClient } from "./api/api";
import { Configuration } from "./api/client";

const graphqlUri: string = process.env.VUE_APP_GRAPHQL_URI
  ? process.env.VUE_APP_GRAPHQL_URI
  : "localhost";
const graphqlPort: string = process.env.VUE_APP_GRAPHQL_PORT
  ? process.env.VUE_APP_GRAPHQL_PORT
  : "9092";
const authMiddleware = new ApolloLink((operation, forward) => {
  const token = localStorage.getItem("user-token");
  operation.setContext({
    headers: {
      authorization: token ? `Bearer ${token}` : null,
    },
  });
  return forward(operation);
});

const httpLink = createHttpLink({
  uri: `https://${graphqlUri}:${graphqlPort}/graphql`,
});

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: authMiddleware.concat(httpLink),
  cache,
});

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
  .provide(DefaultApolloClient, apolloClient)
  .provide(stateSymbol, createState())
  .use(router)
  .mount("#app");
