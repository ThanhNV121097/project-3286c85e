import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        bg: "var(--color-bg)",
        text: "var(--color-text)"
      },
      spacing: {
        6: "var(--space-6)"
      }
    }
  },
  plugins: []
};

export default config;
