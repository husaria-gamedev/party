/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'orange': '#ea5600',
      },
    }
  },
  daisyui: {
    themes: ["business"],
  },
  plugins: [require('daisyui')],
}

