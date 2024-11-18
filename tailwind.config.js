/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{go, js, templ, html}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["black", "lofi"],
  },
};
