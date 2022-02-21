// eslint-disable-next-line @typescript-eslint/no-var-requires
const colors = require('tailwindcss/colors');

module.exports = {
    mode: 'jit',
    content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}', './public/**/*.html', './public/index.html'],
    safelist: [
        {
            pattern: /dark/
        },
        {
            pattern: /primary/
        },
        {
            pattern: /secondary/
        },
        {
            pattern: /info/
        },
        {
            pattern: /danger/
        },
        {
            pattern: /warning/
        },
        {
            pattern: /^border-/,
            variants: ['focus']
        },
    ],
    darkMode: 'class',
    theme: {
        extend: {
            colors: {
                primary: {
                    DEFAULT: '#59344F',
                    '50': '#fbfbfa',
                    '100': '#f4f1f5',
                    '200': '#e8d9ea',
                    '300': '#cdb3d0',
                    '400': '#b387af',
                    '500': '#986390',
                    '600': '#7b4770',
                    '700': '#59344f',
                    '800': '#3e2437',
                    '900': '#241620',
                },
                secondary: {
                    DEFAULT: '#3a4f41',
                    '50': '#f7f9f7',
                    '100': '#ebf1ef',
                    '200': '#d2e1db',
                    '300': '#a4c1b4',
                    '400': '#6b9b87',
                    '500': '#4f7b60',
                    '600': '#3f6146',
                    '700': '#3a4f41',
                    '800': '#233127',
                    '900': '#161e1b',
                },
                info: {
                    DEFAULT: '#224d79',
                    '50': '#f3f8f9',
                    '100': '#d8f1fa',
                    '200': '#ace2f3',
                    '300': '#77c4e2',
                    '400': '#3fa1cb',
                    '500': '#0e79b2',
                    '600': '#27679a',
                    '700': '#224d79',
                    '800': '#183456',
                    '900': '#0f2039',
                },
                danger: {
                    '50': '#fdfcfa',
                    '100': '#fbf0ec',
                    '200': '#f8ced8',
                    '300': '#efa1b3',
                    '400': '#ec718a',
                    '500': '#db5461',
                    '600': '#ca334a',
                    '700': '#a22736',
                    '800': '#761b23',
                    '900': '#491113',
                },
                warning: {
                    DEFAULT: '5b440a',
                    '50': '#fbf9f2',
                    '100': '#f8efad',
                    '200': '#f0de6d',
                    '300': '#d8bb40',
                    '400': '#b49220',
                    '500': '#93740f',
                    '600': '#775b0a',
                    '700': '#5b440a',
                    '800': '#3e2f09',
                    '900': '#2b1d07',
                },
            }
        }
    },
    plugins: [],
}