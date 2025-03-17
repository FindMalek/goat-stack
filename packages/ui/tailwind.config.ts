import type { Config } from "tailwindcss";

const config: Pick<Config, "prefix" | "presets" | "content"> = {
  content: ["./src/**/*.tsx",],
  prefix: "ui-",
  presets: [],
};

export default config;
