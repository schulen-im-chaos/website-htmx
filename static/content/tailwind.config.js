/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./../../web/template/*.{go,templ}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      "light",
      {
        mytheme: {
          "color-scheme": "dark",
          primary: "oklch(65.69% 0.196 275.75)",
          secondary: "oklch(74.8% 0.26 342.55)",
          accent: "oklch(74.51% 0.167 183.61)",
          neutral: "#2a323c",
          "neutral-content": "#e8f0ff",
          "base-100": "#1d232a",
          "base-200": "#191e24",
          "base-300": "#15191e",
          "base-content": "#e8f0ff",
        },
      },
    ],
    darkTheme: "mytheme",
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
