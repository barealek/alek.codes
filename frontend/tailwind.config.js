/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [require('tailwindcss-animated')],
  darkMode: 'class',
  // safelist: [
  //   'animate-delay-75',
  //   'animate-delay-150',
  //   'animate-delay-225',
  //   'animate-delay-300',
  //   'animate-delay-375',
  //   'animate-delay-450',
  //   'animate-delay-525',
  //   'animate-delay-600',
  //   'animate-delay-675',
  //   'animate-delay-750',
  //   'animate-delay-825',
  //   'animate-delay-900',
  //   'animate-delay-975',
  //   'animate-delay-1050',
  //   'animate-delay-1125',
  //   'animate-delay-1200',
  //   'animate-delay-1275',
  //   'animate-delay-1350',
  //   'animate-delay-1425',
  //   'animate-delay-1500',
  //   'animate-delay-1575',
  //   'animate-delay-1650',
  //   'animate-delay-1725',
  //   'animate-delay-1800',
  //   'animate-delay-1875',
  // ]
}

