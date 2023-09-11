/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/*.{html,js}"],
  theme: {
    extend: {
      colors:{
        primary: "#285430",
        second: "#E5D9B6",
      }, 

      screens: {
        '2xl' : '1350px',
      },
    },
  },
  plugins: [],
}

