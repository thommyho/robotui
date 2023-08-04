import { defineSetupVue3 } from "@histoire/plugin-vue";
import "bootstrap/dist/css/bootstrap.min.css";
import smoothscroll from "smoothscroll-polyfill";
import setupI18n from "./js/i18n";
import "./css/app.css";
import { watchThemeChanges } from "./js/theme";

smoothscroll.polyfill();
watchThemeChanges();

export const setupVue3 = defineSetupVue3(({ app }) => {
  app.config.globalProperties.$hiddenFeatures = () => true;
  app.use(setupI18n());
});
