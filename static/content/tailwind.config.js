/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./../../web/template/*.{go,templ}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["light", "dark"],
  },
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui'),
  ]
}