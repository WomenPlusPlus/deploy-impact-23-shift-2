/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,ts}'],
    daisyui: {
        themes: [
            {
                light: {
                    'color-scheme': 'light',
                    'primary-content': '#c9daff',
                    'secondary-content': '#e6d8fb',
                    'accent-content': '#342c5c',
                    'neutral-content': '#D7DDE4',
                    'base-100': '#F6F8FF',
                    'base-200': '#fefdfd',
                    'base-300': '#ffffff',
                    'base-content': '#342c5c',

                    primary: '#4F51B3',
                    secondary: '#e6d8fb',
                    accent: '#f7d2bd',
                    neutral: '#1D1833',
                    info: '#c9daff',
                    success: '#84e9c4',
                    warning: '#fbce60',
                    error: '#f88f8f'
                }
            },
            {
                dark: {
                    'color-scheme': 'dark',
                    'primary-content': '#fefdfd',
                    'secondary-content': '#fefdfd',
                    'accent-content': '#fefdfd',
                    'neutral-focus': '#1d1833',
                    'neutral-content': '#A6ADBB',
                    'base-100': '#1d232a',
                    'base-200': '#191e24',
                    'base-300': '#15191e',
                    'base-content': '#A6ADBB',

                    primary: '#6163bf',
                    secondary: '#d8c2f8',
                    accent: '#f1bb9c',
                    neutral: '#1D1833',
                    info: '#a6c0f7',
                    success: '#84e9c4',
                    warning: '#fbce60',
                    error: '#f88f8f'
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
