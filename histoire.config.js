import { defineConfig } from "histoire";
import { HstVue } from "@histoire/plugin-vue";

export default defineConfig({
  plugins: [HstVue()],
  setupFile: "./histoire.setup.js",
  viteNodeInlineDeps: [/!axios/],
  routerMode: "hash",
  backgroundPresets: [
    {
      label: "default",
      color: "var(--robotui-background)",
      contrastColor: "var(--robotui-default-text)",
    },
    {
      label: "box",
      color: "var(--robotui-box)",
      contrastColor: "var(--robotui-default-text)",
    },
  ],
});
