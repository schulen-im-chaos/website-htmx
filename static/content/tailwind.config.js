/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./../../web/template/*.{go,templ}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          "base-content": "rgb(226 232 240)",
          "neutral-content": "rgb(226 232 240)",
        },
      },
    ],
  },
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui'),
  ]
}