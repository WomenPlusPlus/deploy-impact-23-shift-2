/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,ts}'],
    daisyui: {
        themes: [
            {
                light: {
                    'color-scheme': 'light',
                    'primary-content': '#E0D2FE',
                    'secondary-content': '#FFD1F4',
                    'accent-content': '#07312D',
                    'neutral-content': '#D7DDE4',
                    'base-100': '#F6F8FF',
                    'base-200': '#fafdff',
                    'base-300': '#ffffff',
                    'base-content': '#1f2937',

                    primary: '#4F51B3',
                    secondary: '#8FAFF5',
                    accent: '#ED905C',
                    neutral: '#1D1833',
                    info: '#5CB9ED',
                    success: '#36d399',
                    warning: '#fbbd23',
                    error: '#ED5C5C'
                }
            },
            {
                dark: {
                    'color-scheme': 'dark',
                    'primary-content': '#ffffff',
                    'secondary-content': '#ffffff',
                    'accent-content': '#ffffff',
                    'neutral-focus': '#242b33',
                    'neutral-content': '#A6ADBB',
                    'base-100': '#1d232a',
                    'base-200': '#191e24',
                    'base-300': '#15191e',
                    'base-content': '#A6ADBB',

                    primary: '#6869a4',
                    secondary: '#8FAFF5',
                    accent: '#ED905C',
                    neutral: '#1D1833',
                    info: '#5CB9ED',
                    success: '#36d399',
                    warning: '#fbbd23',
                    error: '#ED5C5C'
                }
            }
        ], // true: all themes | false: only light + dark | array: specific themes like this ["light", "dark", "cupcake"]
        darkTheme: 'dark', // name of one of the included themes for dark mode
        base: true, // applies background color and foreground color for root element by default
        styled: true, // include daisyUI colors and design decisions for all components
        utils: true, // adds responsive and modifier utility classes
        rtl: false, // rotate style direction from left-to-right to right-to-left. You also need to add dir="rtl" to your html tag and install `tailwindcss-flip` plugin for Tailwind CSS.
        prefix: '', // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
        logs: true // Shows info about daisyUI version and used config in the console when building your CSS
    },
    plugins: [require('@tailwindcss/typography'), require('daisyui')]
};
