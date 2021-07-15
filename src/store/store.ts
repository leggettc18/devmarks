import { reactive, provide, inject } from "vue";
import { User } from "@/models/user";
import { AuthState } from "@/models/auth";

export const stateSymbol = Symbol("state");

class State {
  private user: User | null = null;
  private auth: AuthState | null = null;
  private dark = false;

  constructor() {
    const userJson = localStorage.getItem("user");
    if (userJson) {
      this.user = JSON.parse(userJson) as User;
    }
    const token = localStorage.getItem("user-token");
    if (token) {
      this.auth = { token };
    }
    if (
      localStorage.theme === "dark" ||
      (!("theme" in localStorage) &&
        window.matchMedia("(prefers-color-scheme: dark)").matches)
    ) {
      this.dark = true;
    } else {
      this.dark = false;
    }
  }

  storeUser(user: User) {
    this.user = user;
    localStorage.setItem("user", JSON.stringify(user));
  }

  getToken() {
    return this.auth?.token;
  }

  storeToken(authState: AuthState) {
    this.auth = authState;
    localStorage.setItem("user-token", authState.token);
  }

  logOut() {
    this.user = null;
    this.auth = null;
    localStorage.removeItem("user");
    localStorage.removeItem("user-token");
  }

  isUserSet(): boolean {
    return this.user !== null;
  }

  toggleDarkmode() {
    this.dark = !this.dark;
    if (this.dark) {
      localStorage.theme = "dark";
    } else {
      localStorage.theme = "light";
    }
  }

  setDarkmode(darkmode: boolean) {
    this.dark = darkmode;
    if (darkmode) {
      localStorage.theme = "dark";
    } else {
      localStorage.theme = "light";
    }
  }

  darkmodeOSPreference() {
    localStorage.removeItem("theme");
    this.dark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  }

  isDarkmode() {
    return "theme" in localStorage
      ? this.dark
      : window.matchMedia("(prefers-color-scheme: dark)").matches;
  }
}

export const createState = () => {
  return reactive(new State());
};

export const useState = () => inject(stateSymbol) as State;
export const provideState = () => provide(stateSymbol, createState());
